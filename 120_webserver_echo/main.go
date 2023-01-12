package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/HWJskt/learngo/120_webserver_echo/scrapper"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Webserver using Echo framework")
	// scrapper.Scrape("term")

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

// Handler
func handleHome(c echo.Context) error {
	return c.File("home.html")
}

const fileName = "books.csv"

func handleScrape(c echo.Context) error {
	defer os.RemoveAll(fileName)
	fmt.Println(c.FormValue("term"))
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)

	//books.csv라는 파일을 만들고, "책목록"이라는 이름으로 다운로드받는데,
	//books.csv는 중복되니까 마지막에 defer로 삭제했다.
	return c.Attachment("books.csv", "책목록.csv")
}
