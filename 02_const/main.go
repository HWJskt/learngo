/*상수
상수는 한번 선언하면 바꿀수 없다
코드값으로 사용한다
상수는 좌변으로 사용할수 없다 = 메모리 공간이 없다

*/

package main

import "fmt"

func main() {
	//상수는 타입을 안적어도 된다
	//타입을 넣지 않으면 사용할때 타입이 결정이 된다
	const pi = 3.14
	const pii float64 = 3.14
	fmt.Println(pi, pii)

	//iota : 상수 선언때 편하게 하는 방법, 열거값을 편하게 사용하는 방법
	const (
		red   int = iota
		blue  int = iota
		green int = iota
	)
	fmt.Println(red, blue, green)

	const (
		c1 int = iota + 2 //맨 앞에것만 타입, 규칙 써줘도 된다
		c2
		c3
	)
	fmt.Println(c1, c2, c3)

}
