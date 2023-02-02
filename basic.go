package main

import "fmt"

var isActive bool                   // 全域性變數宣告
var enabled, disabled = true, false // 忽略型別的宣告

func main() {
	var x, y int
	x, y = 8, 7

	const Pi float32 = 3.1415926

	fmt.Println("Hello Go", 1+2)
	fmt.Println(x+y, Pi)
}
