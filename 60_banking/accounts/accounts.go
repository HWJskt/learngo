package accounts

import (
	"errors"
	"fmt"
)

// 다른파일에서 접근하려면 type이름 첫글자도 대문자, field이름 첫글자도 대문자여야 한다.
type PublicAccount struct {
	Owner   string
	Balance int
}

// 이름은 Public인데, field는  priviate 이다
type PriviateFiled struct {
	owner   string
	balance int
}

// func를 만들어서 object를 return한다.
// &변수 : 메모리 주소를 의미. *변수 : 메모리주소에 들어있는 값을 의미
// NewwAccount : Account 를 만든다
func NewAccount(owner string) *PriviateFiled {
	account := PriviateFiled{owner: owner, balance: 0}
	return &account
}

// metohd 는 func글자와 func이름 사이에 쓰면 만들어진다.
// p 와  p의 type(struct이름) 을 써주었다
// 이 method 는 p 라는 rceiver를 가진다, p의 type은 PriviateFiled 이다
// receiver의 이름은 strut의 첫글자를 따서 소문자로 작성해야 한다
// p의 type 으로 PriviateFiled 라고 쓰면 p가 계속 복사될뿐 변경되지 않는다.
// * 를 붙이면 Deposit를 호출할때 account 나 receiver를 복사하지 않고 실제 receiver에게 주라는 뜻
func (p *PriviateFiled) Deposit(amount int) {
	p.balance += amount
}

// Balance of tour account
func (p PriviateFiled) Balance() int {
	return p.balance
}

// Error handling

// err로 시작하는 error 내용을 만든다.
var errNoMoney = errors.New("can't withdraw money")

// Withdraw 가 음수가 되지 않도록 error message를 만들자
func (p *PriviateFiled) Withdraw(amount int) error {
	if p.balance < amount {
		// return errors.New("Cant rethdraw your are poor")
		return errNoMoney
	}
	p.balance -= amount
	return nil
}

// owner를 바꾸는 method
func (p *PriviateFiled) ChangeOwner(newOwner string) {
	p.owner = newOwner
}

// owner 를 표시해주는 method
func (p *PriviateFiled) Owner() string {
	return p.owner
}

// String 이라는 이름의 method
// python의 __init__같은 기능, 호출을 하면 나오는 메세지를 지정할 수 있다.
func (p *PriviateFiled) String() string {
	return fmt.Sprint(p.Owner(), "'s account. \n$: ", p.Balance())
}
