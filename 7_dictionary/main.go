package main

import (
	"fmt"

	"github.com/HWJskt/learngo/7_dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	fmt.Println(dictionary)
	fmt.Println(dictionary["first"])

	definition, err := dictionary.Search("Second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	definition2, err2 := dictionary.Search("first")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition2)
	}

	fmt.Println("---------------------------------------------------------")

	// 단어를 더하고, 오루가 나면 오류를 출력한다
	err_add1 := dictionary.Add("hello", "Greeting")
	if err_add1 != nil {
		fmt.Println((err_add1))
	}
	// 위에서 추가한 결과를 확인
	result, _ := dictionary.Search("hello")
	fmt.Println("found", result)

	// 단어를 더하고, 오류가 나면 오류를 출력(이미 단어가 있어서 오류가 출력
	err_add2 := dictionary.Add("hello", "Greeting")
	if err_add2 != nil {
		fmt.Println((err_add2))
	}

	fmt.Println("-----Update---------------------------------------------------")
	// 딕셔너리를 만들고
	dictionary3 := mydict.Dictionary{}
	baseWord3 := "hello world"
	dictionary3.Add(baseWord3, "First")
	fmt.Println(dictionary3.Search(baseWord3))

	// 단어를 업데이트 해본다 (정상))
	err3 := dictionary3.Update(baseWord3, "Second")
	if err3 != nil {
		fmt.Println(err3)
	}
	result3, _ := dictionary3.Search(baseWord3)
	fmt.Println("New result3: ", result3)

	//없는 단어를 업데이트 해본다
	err31 := dictionary3.Update("ㅁㄴㅇㄹ", "Second")
	if err31 != nil {
		fmt.Println(err31)
	}

	fmt.Println("-----Delete--------------------------------------------------")
	dictionary4 := mydict.Dictionary{}
	baseword4 := "hello"
	dictionary4.Add(baseword4, "delete this")
	word41, _ := dictionary4.Search(baseword4)
	fmt.Println(word41)

	del_result4 := dictionary4.Delete(baseword4)
	fmt.Println(del_result4, "<--- nil나오면 지워진거다")

}
