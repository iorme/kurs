package kurs

import (
	"fmt"
	"testing"
	"encoding/json"
	)

func TestParse(t *testing.T) {
	url := "http://www.bankmandiri.co.id/resource/kurs.asp?row=2"
	
	doc := processUrl(url)

	str := getLastUpdated(doc)
	if str == "" {
		t.Error("error get last updated date : ")
		return
	}

	data,_ := parseMandiriHtml(doc)
	if data == nil {
		t.Error("empty data currency")
	}

	response := Data{
		LastUpdated: str,
		Currency: data,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		t.Error("error json marshal")
	}
	fmt.Println(string(jsonResponse))

	doc.Free()
}