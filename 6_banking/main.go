package main

import (
	"fmt"

	"github.com/HWJskt/learngo/6_banking/banking"
)

func main() {
	account := banking.Account{Owner: "nico", Balance: 1000}
	fmt.Println(account)
}
