// 패키지를 컴파일 하고 싶다면 파일 이름 main.go 로 해야 한다.
// folder 안의 모든 파일 중  main function은 하나만 있어야 한다.
// 파일 이름을 다른걸로 하면 기능묶음 정도의 기능을 한다

// 함수
// 함수의 위치는 사용하는 곳 앞,뒤 모두 가능하다

package main // 작성할 패키지 이름을 작성한다

import (
	"fmt"
	"strings"
)

// func 인자, 반환값 모두 type 을 적어줘야한다
func multiply(a int, b int) int {
	return a * b
}

// 인자의 type 을 한번에 쓸수도 있다.
func multiply2(a, b int) int {
	return a * b
}

// func에 반환값이 여러개가 될수있다(go 의 특징)
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// func에 여러개의 인자를 전달하는법
// 반환값이 없고 함수를 실행시키는거니까 반환값 type을 안넣는다
func repeatMe(words ...string) {
	fmt.Println((words))
}

// naked return : 출력값에 이름을 지정할 수 있다.
// 첫줄에 반환변수이름,type을 써주면, return에서 반환변수를 안써도 된다.
func lenAndUpper_naked(name string) (length int, uppercase string) {
	length = len(name) // 첫줄에서 type지정했으니 여기선 안해도된다.
	uppercase = strings.ToUpper(name)
	return // return length, uppercase 이렇게 써도 된다.
}

// defer : func 이 끝나고 나서 실행할 일을 지시할수 있다
func lenAndUpper_defer(name string) (int, string) {
	defer fmt.Println(("I`m done")) // defer 는 func이 값을 return하고 나서 실행된다.
	length := len(name)
	uppercase := strings.ToUpper(name)
	return length, uppercase // return length, uppercase 이렇게 써도 된다.
}

// 컴파일러는 main package 를 찾고, 그 안에 있는 main func 를 찾아서 실행시키다
func main() {

	fmt.Println(multiply(2, 2))
	fmt.Println(multiply2(2, 2))

	totalLength, upperName := lenAndUpper("nicolas")
	fmt.Println(totalLength, upperName)
	totalLength2, _ := lenAndUpper("nicolas") //return 값이 2개인 함수에서 하나만 받고싶으면 _ 사용
	fmt.Println(totalLength2)

	repeatMe("nico", "lynn", "dal", "mark")

	totalLength2, up := lenAndUpper_naked("christmans")
	fmt.Println(totalLength2, up)

	deferlenght, deferuppercase := lenAndUpper_defer("deferfunction")
	fmt.Println(deferlenght, deferuppercase)
}
