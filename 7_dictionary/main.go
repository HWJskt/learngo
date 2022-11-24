package main

import (
	"fmt"

	"github.com/HWJskt/learngo/7_dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "world"
	fmt.Println(dictionary)

}
