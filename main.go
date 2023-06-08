package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/saintfish/chardet"
)

func main() {
	// Start HTTP server
	r := gin.Default()
	r.GET("/scrape", scrapeText)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func scrapeText(c *gin.Context) {
	url := c.Query("url")

	// Getリクエストでレスポンス取得
	res, _ := http.Get(url)
	defer res.Body.Close()

	// Body内を読み取り
	buffer, _ := ioutil.ReadAll(res.Body)

	// 文字コードを判定
	detector := chardet.NewTextDetector()
	detectResult, _ := detector.DetectBest(buffer)
	fmt.Println(detectResult.Charset)
	// => UTF-8

	// 文字コードの変換
	bufferReader := bytes.NewReader(buffer)
	reader, _ := charset.NewReaderLabel(detectResult.Charset, bufferReader)

	// HTMLをパース
	document, _ := goquery.NewDocumentFromReader(reader)

	// titleを抜き出し
	// result := document.Find("山陰合同銀行").Text()
	all_text := document.Text()
	all_text_list := strings.Split(all_text, "\n")
	for _, v := range all_text_list {
		if strings.Contains(v, "山陰合同銀行") {
			fmt.Println(v)
			c.IndentedJSON(http.StatusOK, v)
		}
	}
	// fmt.Println(result)

	// Response: JSON
	// c.IndentedJSON(http.StatusOK, result)
}
