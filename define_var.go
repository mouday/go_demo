package main

import (
	"fmt"
	"reflect"
)

// 全局变量定义，一行声明并赋值多个变量, 不能省var
var a, b, c int32 = 1, 2, 3

var e int  // 声明
var f int = 456  // 声明并赋值

// 分组赋值
var(
	g string = "go语言"
	h int = 1978
)

func main() {
	//局部变量定义, 自动类型推断，简写方式
	d, _, e := 4, 5, 6  // _被丢弃
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	/*
	1
	2
	3
	4
	6
	*/
	
	// 类型转换, 不同类型不能转b
	var f  = float32(a)
	fmt.Println(reflect.TypeOf(f))
	// float32


}
