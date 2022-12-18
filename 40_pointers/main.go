package main

import "fmt"

func main() {
	a := 2
	b := 4
	fmt.Println(&a, &b) //메모리 주소를 보려면 변수명 앞에 &를 붙이면 된다

	c := &a // c는 a의 메모리주소
	fmt.Println(a, &a, c)

	// c는 a의 메모리주소인데 값을 알고싶으면 * 붙인다
	fmt.Println(c, *c)

	//아주 무거운 데이터값들을 다루고 있을때, 값이 같은 여러변수를 두는것보다, 주소를 공유하면 메모리를 절약할수있다
	// c는 a의 메모리주소를 공유하고 있으므로, c의 값을 바꾸면 a 값도 함께 바뀐다
	*c = 20
	fmt.Println(a, *c)
}
