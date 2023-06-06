package main

import "fmt"

func main() {
	// 布尔类型
	var isTrue bool = true
	fmt.Println("布尔类型:", isTrue)

	// 整数类型
	var num1 int = 10
	fmt.Println("整数类型:", num1)

	// 浮点数类型
	var num2 float64 = 3.14
	fmt.Println("浮点数类型:", num2)

	// 字符串类型
	var str string = "Hello, Go!"
	fmt.Println("字符串类型:", str)

	// 数组类型
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("数组类型:", arr)

	// 切片类型
	var slice []int = []int{4, 5, 6}
	fmt.Println("切片类型:", slice)

	// 映射类型
	var m map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("映射类型:", m)

	// 结构体类型
	type Person struct {
		Name string
		Age  int
	}
	var p Person = Person{Name: "Alice", Age: 25}
	fmt.Println("结构体类型:", p)

	// 接口类型
	type Animal interface {
		Sound() string
	}
	type Dog struct{}
	func (d Dog) Sound() string {
		return "Woof!"
	}
	var animal Animal = Dog{}
	fmt.Println("接口类型:", animal.Sound())
}
