// 패키지를 컴파일 하고 싶다면 파일 이름 main.go 로 해야 한다.
// folder 안의 모든 파일 중  main function은 하나만 있어야 한다.
// 파일 이름을 다른걸로 하면 기능묶음 정도의 기능을 한다

// 작성할 패키지 이름을 작성한다
package main

import (
	"fmt"

	"github.com/HWJskt/learngo/something"
)

func main() {
	fmt.Println("hello world")

	//const
	const name string = "const : non changeable"
	// name = "Lynn"  // const 는 바꿀수 없다.
	fmt.Println(name)

	//variable
	var name_var string = "nic"
	name_var = "var : changeable"
	fmt.Println(name_var)

	//type 을 생략하기 (type이 자동으로 지정된다, 지정된 type바꿀수없다. func밖에서는 쓸수없다)
	name_notype := "notype nico"
	name_notype = ":= varable, changeable. Type is fixed. only in func"
	fmt.Println(name_notype)

	//package, import
	//다른 파일에서 기능불러오기 (대문자로 시작하는 함수이름만 불러올수있다)
	something.SayHi()
	// something.sayGoodbye() // 소문자로 시작하는거라 불러올수없다

}
