package something

import "fmt"

// func 이름이 대문자로 시작하면 다른 파일에서 불러와서 쓸수있다. (소문자로 시작하면 안됨)
func SayHi() {
	fmt.Println("func 이름이 대문자로 시작하면 다른 파일에서 불러와서 쓸수있다.")
	fmt.Println("package이름.func이름 형태로 사용 : funcExercise.SayHi()")
}

func sayGoodbye() {
	fmt.Println("소문자로 시작하는 func은 Export 되지않는다 = 다른 파일에서 불러와서 쓸수없다.")
}
