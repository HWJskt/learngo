package main

import "fmt"

func main() {
	// array(베열) : python의 list 같은거, 갯수제한 지정된다
	// array 만드는법 : [최대갯수]tyoe{값들}
	names_array := [5]string{"nico", "lynn", "dal"}
	fmt.Println(names_array)
	// array에 값 추가하기
	names_array[3] = "alala"
	names_array[4] = "bababa"
	fmt.Println(names_array)

	//slice(슬라이스) : 갯수제한이 없는 array, 동적배열
	names_slice := []string{"nico", "lynn", "dal"}
	fmt.Println(names_slice)
	//slice에 값추가하기 : append (slice이름, 값) , 원래 slice를 바꾸지 않는다.
	names_slice = append(names_slice, "flynn")
	fmt.Println(names_slice)

	//map : 파이썬 딕셔너리 처럼 key:value 형태
	//mpa 만드는법 : map[key tyep]value type{값들}
	nico_map := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico_map)
	//map에서 loop 사용가능
	for key, value := range nico_map {
		fmt.Println(key, value)
	}

	//Structs : go 에서  JSON형태의 자료를 넣는거
	//Structs 방법1
	//go에는 python의 __init__ 같은 constructor 가 없다
	type person struct {
		name    string
		age     int
		favFood []string
	}
	myFood := []string{"kimchi", "ramen"}
	nico := person{"nico", 18, myFood}
	//Structs 방법1
	flynn := person{name: "flynn", age: 17, favFood: myFood} // field:value 형태  // 이 코드가 더 명확해보인다.
	//Structs 출력
	fmt.Println(nico)
	fmt.Println(nico.name)
	fmt.Println(flynn)

}
