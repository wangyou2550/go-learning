package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 在临界区进行操作
	c.count++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 在临界区进行操作
	return c.count
}

func main() {
	counter := Counter{}

	// 启动多个 goroutine 并发地增加计数器的值
	for i := 0; i < 10; i++ {
		go func() {
			counter.Increment()
		}()
	}

	// 等待一段时间以确保所有 goroutine 执行完成
	time.Sleep(time.Second)

	// 获取最终的计数器值并打印
	fmt.Println(counter.GetValue())
}
