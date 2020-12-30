package main

import (
	"fmt"
	"log"
	"net/http"
)

var baseURL string = "https://airbnb.com"

func main() {
	pages := getPage()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	//goquery document 만들기
	doc, err := goquery.NewDocumentFrontReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each()

	fmt.Println(doc)

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status", res.StatusCode)
	}
}
