package mydict

import "errors"

// Dictionary type 이다. struct type 이 아니다
type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errWordExists = errors.New("that word already exists")

// Search : type은 method를 가질수 있다.
// map의 key 를 호출하면 2개의 값을 얻는다. value, exists(이건 boolean이다)
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	// Dictionary에서 Search 기능을 사용 - 단더가 있으면 err 가  nill, 없으면  errNotFound 반환한다
	_, err := d.Search((word))
	//방법1
	// if err == errNotFound {
	// 	d[word] = def
	// } else {
	// 	return errWordExists
	// }

	//방법2
	switch err {
	case errNotFound:
		d[word] = def
		// return nil 로 이어진다
	case nil:
		return errWordExists
	}
	return nil
}
