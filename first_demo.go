package main // 包名称，只能在文件开头

import "fmt"  // 导入包

const NAME  string= "peng shiyu"  // 常量

var global string= "global"   // 全局变量

// 一般类型声明，类型别名
type i int

// 声明结构体
type Learn struct {

}

// 声明接口
type Ilearn interface {

}
/*
多行注释
*/

// 自定义函数
func myfunc() {
	fmt.Println("myfunc")
}

// main函数，入口函数
func main() {
	fmt.Println("hello world")
	fmt.Println(NAME)
	fmt.Println(global)
	myfunc()
}
