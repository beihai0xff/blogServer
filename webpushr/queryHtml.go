package webpushr

import (
	"os"

	"github.com/PuerkitoBio/goquery"
)

type pageInfo struct {
	title string
	url   string
}

// 获取文章 title 和 url
func query(filepath string) (pageInfo, error) {
	var p pageInfo
	err := PathExists(filepath)
	if err != nil {
		return p, err
	}
	res, err := os.Open(filepath)
	if err != nil {
		return p, err
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		return p, err
	}
	// Find the review items
	dHead := doc.Find("head")

	p.title, _ = dHead.Find("meta[name=\"twitter:title\"]").Attr("content")
	p.url, _ = dHead.Find("meta[property=\"og:url\"]").Attr("content")
	return p, err
}
