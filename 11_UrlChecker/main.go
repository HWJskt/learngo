package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	// empty map 을 만드는법
	// var results map[string]string  // 이 map 은 nill 이 되어 이건 나중에 값을 넣을때 오류난다
	var results = map[string]string{}
	results["hello"] = "hello"

	c := make(chan requestResult)

	// urls 는 slice 이다
	urls := []string{
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.naver.com/",
		"https://hwjskt.github.io/",
	}

	// 인덱스를 _ 처리하고 value를 출력한다
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

// url을 클릭하면 결과가 어떻게나오는지 출력한다
// golang std library 에서 net / http 모듈사용법을 확인할수있다
func hitURL(url string, c chan<- requestResult) {
	fmt.Println("checking: ", url)
	resp, err := http.Get(url)
	var status = "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
