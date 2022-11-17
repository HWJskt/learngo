package main

import (
	"fmt"

	"github.com/HWJskt/learngo/6_banking/accounts"
)

func main() {
	// 접근가능한 const : 이름과 field가 대문자로 정의되어 정의, 수정 모두 된다.
	publicAccount := accounts.PublicAccount{Owner: "nico", Balance: 1000}
	publicAccount.Owner = "lynn" // 계좌 주인이 바뀌어 버렸다
	fmt.Println(publicAccount)

	account := accounts.NewAccount("nico")
	// account.balance = 1000 //이거 안된다. 왜냐하면 balance는 account.go에서 func정의할때 priviate로 만들었으니까
	fmt.Println(account)

}
