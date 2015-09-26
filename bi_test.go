package kurs

import (
	"fmt"
	"testing"
	"encoding/json"
	)

func TestParseBi(t *testing.T) {
	url := "http://www.bi.go.id/id/moneter/informasi-kurs/transaksi-bi/Default.aspx"
	
	doc := processUrl(url)

	str := getLastUpdatedBi(doc)
	if str == "" {
		t.Error("error get last updated date : ")
		return
	}

	data,_ := parseBiHtml(doc)
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