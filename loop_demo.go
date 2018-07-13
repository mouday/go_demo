package main

import (
	"fmt"
	"time"
)

//if 的用法
func if_func(){
	var a = 0
	if a ==0 {
		fmt.Println("a == 0")
	} else {
		fmt.Println("a != 0")
	}
}

// switch
func switch_func(){
	var b interface{}
	b = 32
	switch b.(type){
	case int:
		fmt.Println("a type is int")
	case string:
		fmt.Println("a type is string")
	default:
		fmt.Println("以上不满足")
	}
}

//无限循环
func while_func(){
	for {
		fmt.Println("无限循环")
	}
}


//条件循环
func for_func(){
	for i:=1; i<=10; i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// 遍历数组
func for_list(){
	var list = []string{"苹果", "香蕉", "菠萝"}
	for _, value:= range list{
		fmt.Println(value)
	}
}

func goto_demo(){
	goto Label
	fmt.Println("这句话不会执行")
	Label:
		fmt.Println("直接执行这句话")
}

func main() {
	goto_demo()
}