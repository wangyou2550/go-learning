// 在上述代码中，我们创建了一个 JobQueue 类型，其中包含一个互斥锁 mu 和一个条件变量 cond。AddJob() 方法用于向作业队列中添加作业，并在添加后通过 cond.Signal() 通知等待的消费者 goroutine。GetJob() 方法用于从作业队列中获取作业，如果队列为空，则通过 cond.Wait() 进行等待，直到有新的作业被添加。

// 在 main() 函数中，我们启动了一个消费者 goroutine，该 goroutine 不断从作业队列中获取作业并进行处理。然后，我们通过调用 AddJob() 方法添加一些作业到队列中，最后等待一段时间以确保消费者有足够的时间处理所有作业。

// 通过使用 sync.Cond，我们可以在生产者和消费者之间实现一种同步机制，以便在作业队列有新作业时及时通知消费者进行处理。这样可以避免消费者长时间无作业可处理的空闲等待，并确保生产者和消费者之间的正确同步
package main

import (
	"fmt"
	"sync"
	"time"
)

type JobQueue struct {
	mu      sync.Mutex
	cond    *sync.Cond
	jobList []string
}

func NewJobQueue() *JobQueue {
	jobQueue := &JobQueue{
		jobList: make([]string, 0),
	}
	jobQueue.cond = sync.NewCond(&jobQueue.mu)
	return jobQueue
}

func (jq *JobQueue) AddJob(job string) {
	jq.mu.Lock()
	defer jq.mu.Unlock()

	jq.jobList = append(jq.jobList, job)
	jq.cond.Signal()
}

func (jq *JobQueue) GetJob() string {
	jq.mu.Lock()
	defer jq.mu.Unlock()

	for len(jq.jobList) == 0 {
		jq.cond.Wait()
	}

	job := jq.jobList[0]
	jq.jobList = jq.jobList[1:]
	return job
}

func main() {
	jobQueue := NewJobQueue()

	// 启动一个消费者 goroutine
	go func() {
		for {
			job := jobQueue.GetJob()
			fmt.Println("Processing job:", job)
		}
	}()

	// 生产者添加任务
	jobQueue.AddJob("Job 1")
	time.Sleep(1 * time.Second)
	jobQueue.AddJob("Job 2")
	time.Sleep(1 * time.Second)
	jobQueue.AddJob("Job 3")

	// 等待一段时间以确保消费者有足够的时间处理所有任务
	time.Sleep(3 * time.Second)
}
