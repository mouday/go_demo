package packageB

import "fmt"

func init() {
	fmt.Println("packageB.fileB.init")
}

var car = "汽车"
var House = "房子"

func funcB()  {
	fmt.Println("packageB.fileB.funcB")
}