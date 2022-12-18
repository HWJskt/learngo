package main

import (
	"fmt"
	"time"
)

//goroutine끼리 정보를 주고 받을 수 있다
//goroutine이 main function으로 메세지를 주고 받을수있다.

func main() {
	//channel을 만드는법 chan이라고 쓰고 type을 지정해줘야한다
	c := make(chan string)

	people := [2]string{"nico", "lynn"}
	for _, person := range people {
		go isSexy(person, c) //채널을 두함수에 보내고 있다
	}

	//채널을 통한 메세지를 받기
	// 채널에서 메서지를 받을때까지 기다린다. 그럼 goroutine이 계속 작동한다.
	// <-c : 메세지를 받는다는것은 blocking operation이다

	// resultOne := <-c
	// resultTwo := <-c
	// fmt.Println(resultOne)
	// fmt.Println(resultTwo)

	// <-c 통해서 받는 메세지는 goroutine갯수와 같아야 한다.
	// fmt.Println("received this message:", <-c)
	// fmt.Println("received this message:", <-c) // 함수가 2개면 2번 메세지를 받아야 한다

	//loop를 이용해서 채널 메세지를 받을수도 있다
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

// func에 채널 만드는법
// 두번째 인자에 채널이름, chan, 채널데이터 타입 을 추가했다
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 3)
	c <- person + " is sexy" //35초후에 채널에 이 메세지 보낸다
}
