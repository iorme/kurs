package kurs

import (
	"net/http"
	"io/ioutil"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/html"
	)

func processUrl(url string) *html.HtmlDocument {
	respon, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	page, _ := ioutil.ReadAll(respon.Body)
	document, _ := gokogiri.ParseHtml(page)

	return document
}