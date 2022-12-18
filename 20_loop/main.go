/* go에서는 반복문이 for 뿐이다

for 초기문; 조건문; 후처리; {
	코드블록
}

조건문이 false가 될때까지 실행된다. false가 되면 종료

*/

package main

import "fmt"

// loop : for, range 밖에 없다.
func loopExample(numbers ...int) int {
	//인자로 받은 숫자들은 array 형태로 저장된다.
	fmt.Println(numbers)

	// range는 for 안에서만 사용할 수 있다.
	// 이렇게 하면 index만 나온다
	for number := range numbers {
		fmt.Println(number)
	}

	// 이렇게 해야 숫자를 출력할 수 있다.
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	// for i 를 사용하는 방법
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func superAdd(numbers ...int) int {
	total := 0 //이거 빼먹기 쉽다

	// _ 로 index value 를 ignore한다
	for _, number := range numbers {
		total = total + number
	}

	return total //이거 없으면 안된다
}

func main() {
	loopExample(11, 12, 13, 14, 15)
	result := superAdd(11, 12, 13, 14, 15)
	fmt.Println(result)
}
