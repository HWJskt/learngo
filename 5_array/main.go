package main

import "fmt"

func main() {
	// array
	// array 만드는법 : [최대갯수]tyoe{값들}
	names_array := [5]string{"nico", "lynn", "dal"}
	fmt.Println(names_array)
	// array에 값 추가하기
	names_array[3] = "alala"
	names_array[4] = "bababa"
	fmt.Println(names_array)

	//slice : 갯수제한이 없는 array
	names_slice := []string{"nico", "lynn", "dal"}
	fmt.Println(names_slice)
	//slice에 값추가하기 : append (slice이름, 값) , 원래 slice를 바꾸지 않는다.
	names_slice = append(names_slice, "flynn")
	fmt.Println(names_slice)

	//map
	//mpa 만드는법 : map[key tyep]value type{값들} , 값들은 파이썬 딕셔너리 처럼 key:value 형태
	nico_map := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico_map)
	//map에서 loop 사용가능
	for key, value := range nico_map {
		fmt.Println(key, value)
	}
}
