package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico", 10)
	sexyCount("flynn", 4)
	//goroutine 은 main function이 실행되는 중에만 실행되는거다
	//->두번째 함수에 go 를 추가하면 프로그램이 종료되어 버린다.
	//두번째 함수에 go 가 없어야 두번째 함수가 실행되는 동안 첫번째함수(goroutine적용된)가 작동하는것
	//두번째 함수가 종료되면 goroutine 이 종료되어 첫번째 함수도 종료된다.

}

func sexyCount(person string, t int) {
	for i := 0; i < t; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
