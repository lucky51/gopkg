package crawler

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var pkgSearchUrl = "https://pkg.go.dev/search?q=%s"

type PkgS struct {
	PkgName        string
	PkgPath        string
	PkgDescription string
}

func saveFile(fileName string, data []byte) error {
	return ioutil.WriteFile(fileName, data, 0777)
}
func Crawl(keyword string, limit int) ([]PkgS, error) {
	var pkgs = make([]PkgS, 0)
	if keyword == "" {
		return pkgs, nil
	}
	c := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
		},
	}
	resp, err := c.Get(fmt.Sprintf(pkgSearchUrl, keyword))
	if err != nil {
		log.Fatalf("%v", err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err.Error())
	}
	sections := doc.Find("div.SearchSnippet")
	sections.EachWithBreak(func(idx int, sec *goquery.Selection) bool {
		if limit > 0 && limit-1 < idx {
			return false
		}
		var snipTitleNode = sec.Find("div.SearchSnippet-headerContainer>h2>a")
		pkgName := ""
		snipTitleNode.Contents().EachWithBreak(func(idx int, txtSec *goquery.Selection) bool {
			if goquery.NodeName(txtSec) == "#text" {
				pkgName = txtSec.Text()
				pkgName = strings.Replace(pkgName, "\n", "", -1)
				pkgName = strings.TrimSpace(pkgName)
				return false
			}
			return true
		})
		pkgPathNode := snipTitleNode.Find("span.SearchSnippet-header-path")
		pkgPath := pkgPathNode.Text()
		pkgPath = strings.TrimPrefix(pkgPath, "(")
		pkgPath = strings.TrimSuffix(pkgPath, ")")
		pkgDescription := sec.Find("p.SearchSnippet-synopsis").Text()
		pkgDescription = strings.Replace(pkgDescription, "\n", "", -1)
		pkgDescription = strings.TrimSpace(pkgDescription)
		if pkgPath != "" {
			pkgs = append(pkgs, PkgS{
				PkgPath:        pkgPath,
				PkgName:        pkgName,
				PkgDescription: pkgDescription,
			})
		}
		return true
	})
	return pkgs, nil
}
