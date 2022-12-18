/*
# 연산자




*/

package main

import (
	"fmt"
	"math"
)

func main() {

	// 쉬브프 연산
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x: %08b x<<2:%08b x<<2 %d\n", x, x<<2, x<<2)
	fmt.Printf("y: %08b y<<2:%08b y<<2 %d\n", y, y<<2, y<<2) //8자리를 넘어가면 없어진다

	//math.Nextafter()  두값이 있으면 1비트
	fmt.Println("================================================================")
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f == %0.18f : %v\n", a+b, c, equal(a+b, c))

	//논리 연산자
	//	&& : AND  || : OR  ! : NOT
	fmt.Println("================================================================")

	//값 교환
	fmt.Println("값쇼환============================================================")
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

}

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b

}
