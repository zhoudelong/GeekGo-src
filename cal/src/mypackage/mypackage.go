package mypackage

import "fmt"

func New() {
	fmt.Println("这个没有返回...")
}

func MyNew() string {
	fmt.Println("这是我新创建的pkg.....")
	return "这是我新创建的pkg"
}
