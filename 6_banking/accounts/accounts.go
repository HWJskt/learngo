package accounts

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
func NewAccount(owner string) *PriviateFiled {
	account := PriviateFiled{owner: owner, balance: 0}
	return &account
}
