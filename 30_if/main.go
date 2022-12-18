/* 조건문 : if, switch
조건이 여러개면 소괄호를 이용해서 보기좋게 한다



*/

package main

import "fmt"

func canIDraink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

// variable expression : if 조건문에서만 사용하기위해 변수를 만들어 사용
// if 와 조건 사이에 삽입
// if 초기문; 조건문  형태로 사용, 사이는 ; 를 적어준다
func koreanDraink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	// 위, 아래 코드는 같은 뜻이다
	// koreanAge := age + 2
	// if koreanAge < 18 {
	// 	return false
	// }

	return true
}

// switch
//
//	switch 비교찺
//	case 값1: // 비교값이 값1과 같을때 아래 내용을 수행한다
//	default: // 모든 case가 아닐때
//	case 다음에  break를 안써도 된다.
//		(다른언어는 break안쓰면 case를 빠져나가지 않고 다른 case도 사용))
//		아래 case도 사용하고 싶으면 fallthrough 를 써주면 된다
func switchDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func switchCarDrive(age int) bool {
	switch {
	case age <= 18:
		return false
	case age > 18:
		return true
	}
	return false
}

// switch 초기문
func switchKoreanDrive(age int) bool {
	switch korAge := age + 2; korAge {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

// switch 여러개의 값을 비료할 수 있다
func schoolDay(day string) string {
	switch day {
	case "monday", "thuesday":
		return "학교가는날"
	case "wednesday", "thursday", "friday":
		return "학교가는데 실습있는날"
	default:
		return "주말입니다"
	}
}
func main() {
	fmt.Println(canIDraink(16))
	fmt.Println(koreanDraink(16))
	fmt.Println(switchDrink(18))
	fmt.Println(switchCarDrive(18))
	fmt.Println(switchCarDrive(18))
	fmt.Println(switchKoreanDrive(18))
	fmt.Println(schoolDay("monday"))

}
