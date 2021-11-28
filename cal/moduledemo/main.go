package main

import (
	"fmt"
	"moduledemo/modulepkg"
	"src/gitpro" // 调用不同文件目录下的pkg
	"src/multi"
)

func main() {
	fmt.Println("Hello moduledemo....")
	gitpro.Mysqlfunclocal()

	var a, b int
	a, b = 3, 4
	fmt.Println(multi.Add(a, b))

	fmt.Println(modulepkg.Modulepkg())
}
