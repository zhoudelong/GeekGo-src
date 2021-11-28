package main

import (
	"fmt"
	"moduledemo/modulepkg"
	"reflect"
	"src/multi" // 导入自己写的pkg
	"src/mypackage"
)

//引入包问题解决
// 1 go.mod: https://www.liwenzhou.com/posts/Go/import_local_package_in_go_module/
// 2: 通过生成bin https://blog.csdn.net/BTnode/article/details/104763932

/*
go.mod 包管理工具
bin存放编译后的可执行文件；pkg存放编译后的包文件；src存放项目源文件 一般，bin和pkg目录可以不创建
你自己的写的包必须在$GOPATH/src 下
在包中定义的函数首字母必须大写（外部可调用）
 GO 的包搜索是从 GOPATH 和 GOROOT 路径下搜索
源码必须要放在 GOROOT 或 GOPATH 的 src 目录下才能找到

这里有一个Project GOPATH，
还有一个Global GOPATH,把你的项目配置在Project GOPATH里，
每个项目都不一样，创建另一个项目时这个路径要配置成新项目的。
lobal GOPATH可以弄一个公共项目，以后就把第三方的包直接装到这里，就可以自动在你的项目里引用了。

*/

/*
1: GVM（Go Version Manager）多版本golang到本地缓存，通过切换环境变量指向不同的Golang版本
2: Go mod 对 go modules 的依赖管理，使得不同项目中依赖不同版本的同一个库更简单。
 go mod 它可以管理一个依赖库的多个版本同时存在。
*/

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello, golang")

	var a int
	var b int
	a, b = 3, 4
	fmt.Println(Add(a, b))
	fmt.Println(reflect.TypeOf(a))

	// 调用同级目录下的包 multi
	fmt.Println(multi.Multiply(a, b)) // multi 要在goroot/src下,这个包
	fmt.Println(multi.Add(a, b))
	fmt.Println(multi.Subtract(3, 4))

	// 调用同级目录下的包 mypackage
	mypackage.New()
	fmt.Println(mypackage.MyNew())

	// 调用不同目录下的包 moduledemo
	fmt.Println(modulepkg.Modulepkg()) // moduledemo

	////数组切割, 迭代输出
	strings := []string{"google", "runoob", "Hello"}
	fmt.Println(strings, reflect.TypeOf(strings))
	for i, s := range strings {
		fmt.Println(i, s)
	}

}
