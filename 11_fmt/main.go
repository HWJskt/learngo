package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32412.2342352345123124124

	//#표준출력
	// 한줄로 표시, 인자사이에는 띄어쓰기 1칸
	fmt.Println("a:", a, "b", b, "f:", f)

	// formatter : 서식을 지정하고 출력
	// %v 로 쓰면 그냥 알아서 타입맞춰서 출력해준다
	fmt.Printf("a: %d, b: %d, f: %f\n", a, b, f)

	// \n 줄바꿈 \t 탭을 삽입 \\ \출력 \" "를 출력
	fmt.Println(" 줄바꿈\n 탭\t 특부문자출력\\ 특수문자출력\" ")

	//#표준입력
	// scan() : 표준입력에서 값을 입력받는다
	// scanf() : 서식형태로 값을 받는다
	// scanln() : 한줄을 읽어서 값을 입력받는다, 변수이름 안된다. 주소를 넣어주어야한다
	var o int
	var p int
	n, _ := fmt.Scanln(o, p)
	fmt.Println("n,o,p : ", n, o, p)

	m, _ := fmt.Scanln(&o, &p)
	fmt.Println("m,o,p : ", m, o, p)

	// 입력버퍼 : 입력된 내용을 임시로 저장하는 공간
	// scanln 으로 읽고나서 오류가 나면 버퍼값을 비워줘야한다
	fmt.Println("\n입력버퍼================================================================")
	fmt.Println("x,y 값을 int 형태로 띄어쓰기로 구분해서 입력")
	var stdin = bufio.NewReader(os.Stdin) //Stdin:standard input (표준입력) 을 의미

	var x int
	var y int

	_, err := fmt.Scanln(&x, &y) // pointer 주소 형태로 받아야 한다
	if err != nil {              //에러가 있다면
		fmt.Println("에러: ", err)
		stdin.ReadString('\n') //어떤문자가 나올때까지 읽어라 = \n 이 나올때까지 읽어라 = buffer를 비워줘라
	} else {
		fmt.Printf("정상 x=%v y=%v\n", x, y)
	}

}
