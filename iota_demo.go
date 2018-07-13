package main

import "fmt"

const a = iota
const b = iota

const (
	c = iota
	_ = iota  // 跳过1值
	d = iota
	e = 3.14  // 插队使用
	f = iota
	g         // 隐式使用
	h
	j, k = iota, iota + 3  //同一行iota不增加
	l, m
	n = iota
)

func main() {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)
	fmt.Println("h=", h)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	fmt.Println("m=", m)
	fmt.Println("n=", n)
}

/*
a= 0
b= 0
c= 0
d= 2
e= 3.14
f= 4
g= 5
h= 6
j= 7
k= 10
l= 8
m= 11
n= 9
 */