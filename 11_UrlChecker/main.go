package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	// urls 는 slice 이다
	urls := []string{
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.naver.com/",
		"https://hwjskt.github.io/",
	}
	// 인덱스를 _ 처리하고 value를 출력한다
	for _, url := range urls {
		// fmt.Println(url)
		hitURL(url)
	}
}

var errRequestFailed = errors.New("request failed")

// url을 클릭하면 결과가 어떻게나오는지 출력한다
// golang std library 에서 net / http 모듈사용법을 확인할수있다
func hitURL(url string) error {

	resp, err := http.Get(url)
	fmt.Println("Checking: ", url, " --> Result: ", resp.StatusCode)
	// StatusCode 400부터는 오류가 있는 경우
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
