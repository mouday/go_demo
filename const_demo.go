package main

import (
	"fmt"
	"reflect"
)

const NAME string= "姓名"
const AGE = "年龄"
const (
	cat string = "猫"
	dog = "狗"
)

const apple, banana string = "苹果", "香蕉"

func main(){
	fmt.Println(NAME)
	fmt.Println(AGE)
	fmt.Println(cat)
	fmt.Println(dog)
	fmt.Println(apple)
	fmt.Println(banana)
	fmt.Println(reflect.TypeOf(AGE))
	fmt.Println(len(NAME))
}
