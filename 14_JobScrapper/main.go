package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedBook struct {
	title string
	price string
}

var baseURL string = "https://search.shopping.naver.com/book/search?bookTabType=ALL&pageIndex=2&pageSize=40&query=golang&sort=REL"

func main() {
	var books []extractedBook //empty slice를 만들어둔다
	c := make(chan []extractedBook)

	//페이지 수 확인
	totalPages := getPages()

	//각 페이지의 내용을 추출 -> slice를 만들고, books라는 slice로 합쳐서 모아둔다
	//페이지 수 만큼 goroutine을 만든다
	for i := 1; i < totalPages+1; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedBooks := <-c
		books = append(books, extractedBooks...) // ... 는 slice를 합치는 방법
	}

	// slice를 출력
	fmt.Println("")
	fmt.Println(books)
	fmt.Println("=======추출 완료=======")
	// fmt.Println("Total pages: ", totalPages)

	//만든 slice를 모아서 파일로 만드는 함수에 전달
	writeBooks(books)
	fmt.Println("=======csv저장 완료=======")
}

// 페이지 큰 내용(카드)을 추출하는 함수
func getPage(page int, mainC chan<- []extractedBook) {
	var books []extractedBook
	c := make(chan extractedBook)
	pageURL := "https://search.shopping.naver.com/book/search?bookTabType=ALL&pageIndex=" + strconv.Itoa(page) + "&pageSize=40&query=golang&sort=REL"
	fmt.Println("Requesting page : ", page)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".bookListItem_item_book__1yCey")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractBook(card, c)
		// books = append(books, book)
	})

	//카드 갯수만큼 채널 메세지를 받는다 -> 메세지를 slice에 추가
	for i := 0; i < searchCards.Length(); i++ {
		book := <-c                 //message가 오면 받는다
		books = append(books, book) //메세지를 slice에 추가한다.
	}
	mainC <- books
}

// 추출된 카드에서 세부 내용을 찾는 함수 -> 정리된 arrya로 반환한다
func extractBook(card *goquery.Selection, c chan<- extractedBook) {
	title := cleanString(card.Find(".bookListItem_text__bglOw").Text())
	price := cleanString(card.Find(".bookPrice_price__zr5dh").Text())
	// fmt.Println(title, price)
	c <- extractedBook{title: title, price: price}

}

// 섭속해서 페이지 수를 세는 함수
func getPages() int {
	pages := 0
	//접속이 잘되는지 확인한다
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	//getPages함수가 끝났을때 res.Body를 닫아서 메모리 절약
	defer res.Body.Close()

	//웹페이지 내용을 검색한다.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	//class pagination에서 자료를 찾는다
	doc.Find(".Paginator_list_paging__OTDcR").Each(func(i int, s *goquery.Selection) {
		// fmt.Println(s.Html()) //html전체를 보는 코드
		pages = s.Find("a").Length()

	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

// 내요을 csv파일로 기록하는 기능
func writeBooks(books []extractedBook) {
	file, err := os.Create("books.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() //함수가 끝나는 순가 기록한다

	headers := []string{"title", "price"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, book := range books {
		bookSlice := []string{book.title, book.price}
		jwErr := w.Write(bookSlice)
		checkErr(jwErr)
	}
}

// // 글자사이 공백을 제거해주는 함수
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), "")

}
