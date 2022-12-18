/*
# 변수 : 값을 저장하는 메모리 공간을 가리키는 이름
	변수이름 : camel방법을 사용
  		Ex. camelType 소문자로 시작, 두번째 단어첫을자를 대문자로
	변수선언 : 변수키워드 변수이름 타입 어싸인 값
		Ex. var a int = 10
	변수의 4가지 속성 : 이름, 값, 주소, 타입(size가 고정되어있다)

# unsigned int : 부호없는 정수(양수)
	unit8(=byte) 1바이트 부호없는 정수 0~255
	unit16 2바이트 부호없는 정수 0~65535
	unit32 4바이트 부호없는 정수 0~4294967295
	unit64 8바이트 부호없는 정수 0~18446744073709551615

# int : 부호있는 정수
	int8 1바이트 부호있는 정수	-128~127
	int16 2바이트 부호있는 정수
	int32(=rune) 4바이트 부호있는 정수
	int64 8바이트 부호있는 정수

# float
	float32 4바이트 실수 : 소수부 7자리
	float64 8바이트 실수 : 소수부 15자리

byte = uint8 : 1byte 숫자
rune = int32 : 문자1개, UTF-8, 1-3bytes
int : 32비트 컴퓨터에서는 int32, 64비트에서는  int64
unit : 32비트 컴퓨터에서는 uint32, 64비트에서는  uint64

# bool
	true, false 로만 사용가능하다
	1,0 로 사용불가

*/

package main

import (
	"fmt"
)

var g int = 10 //패키지 전역변수. 패키지 내부어디서든 사용가능

func main() {
	// 실수에서 정수르 바꾸면 소수점이 날라간다
	var a float64 = 3.5
	var b int8 = int8(a) //int16 -> int8로 변환
	fmt.Println(a, b)

	// 큰 type에서 작은 type으로 바꾸면 값이 날라간다
	var c int16 = 3456
	var d int8 = int8(c) //int16 -> int8로 변환
	fmt.Println(c, d)

	//변수가 유효한 범위는 변수가 속한 {} 내부 뿐이다
	var m int = 20
	{
		var s int = 50
		fmt.Println("s=", s, ":{}안에서는 사용가능 ")
	}
	m = m + g //정상: g는 전역변수라서 사용가능
	// m = m + s //오류: s는 {}밖에서 없어진다
	fmt.Println("m=", m)

	// 실수연산에서는 자릿수때문에 오차가 생긴다
	var a4 float32 = 1234.523
	var b4 float32 = 3456.123
	var c4 float32 = a4 * b4
	fmt.Println(a4, "*", b4, "!=", c4)

	//type : 새로운 타입을 만들어 쓰겠다
	//type 타입이름 진짜타입
	type myInt int
	var o myInt = 2

	//변수 선언하는 방법
	var p int = 3
	var q int //값을 할당하지 않아 0 이 된다
	var r = 4
	s := 5
	fmt.Println(o, p, q, r, s)

	//타입별 기본값
	//모든 정수타입 : 0
	//모든 실수타입 : 0.0
	//불리언 : false
	//문자열 : "" 빈문자열
	//그외 : nil (정의되지 않은 메모리주소를 나타냄)

}
