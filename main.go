/*
패키지를 컴파일 하고 싶다면 파일 이름 main.go 로 해야 한다.
folder 안의 모든 파일 중  main function은 하나만 있어야 한다.
파일 이름을 다른걸로 하면 기능묶음 정도의 기능을 한다

작성할 패키지 이름을 작성한다
main이라는 package는 프로그램의 시작점을 포함하고 있다는 의미
프로그램은 main package하나와 다른 package들로 이루어져 있다.
*/
package main

// 외부 package를 가져온다
import (
	"fmt"

	"github.com/HWJskt/learngo/something"
)

// main함수는 프로그램 시작점이라는 의미(가장먼저 실행되는 하뭇)
func main() {
	fmt.Println("hello world")

	//const : 변수 - 바꿀수 없다
	const name string = "# const : non changeable"
	// name = "Lynn"  // Error : const 는 바꿀수 없다.
	fmt.Println(name)

	//var : 변수 - 바꿀수 있다(variable)
	var name_var string = "nic"
	name_var = "# var : changeable (from nic)"
	fmt.Println(name_var)

	//type 을 생략하기 (type이 자동으로 지정된다, 지정된 type바꿀수없다. func밖에서는 쓸수없다)
	name_notype := "notype nico"
	name_notype = "# := Type 자동 지정. 이후 Type은 바꿀수 없고, 값은 바꿀수 있다. func 안에서만 쓸수있다"
	fmt.Println(name_notype)

	//package, import
	//다른 파일에서 기능불러오기 (대문자로 시작하는 함수이름만 불러올수있다)
	something.SayHi()
	// something.sayGoodbye() // 소문자로 시작하는거라 불러올수없다

}
