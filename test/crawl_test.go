package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/PuerkitoBio/goquery"
	crawler "github.com/lucky51/gopkg/internal/crawler"
)

func TestCrawlPkg(t *testing.T) {
	_, err := crawler.Crawl("colly")
	if err != nil {
		t.Error(err.Error())
	}
}
func FailureOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func TestCrawlDocQuery(t *testing.T) {
	body, err := ioutil.ReadFile("htmlbody")
	FailureOnError(err)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	FailureOnError(err)
	doc.Find("div.SearchSnippet span.SearchSnippet-header-path").Each(func(idx int, selection *goquery.Selection) {
		fmt.Println(idx, ":", selection.Text())
	})

}
