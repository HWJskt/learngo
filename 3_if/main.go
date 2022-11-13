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

func switchKoreanDrive(age int) bool {
	switch korAge := age + 2; korAge {
	case 10:
		return false
	case 20:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDraink(16))
	fmt.Println(koreanDraink(16))
	fmt.Println(switchDrink(18))
	fmt.Println(switchCarDrive(18))
	fmt.Println(switchCarDrive(18))
	fmt.Println(switchKoreanDrive(18))

}
