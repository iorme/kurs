package kurs

import(
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/moovweb/gokogiri"
)

func processMandiri() Data{
	url := "http://www.bankmandiri.co.id/resource/kurs.asp?row=2"
	
	respon, _ := http.Get(url)
	page, _ := ioutil.ReadAll(respon.Body)
	document, _ := gokogiri.ParseHtml(page)
	defer document.Free()

	nilai := "1.00"
	var matauang string
	var kursjual string
	var kursbeli string
	kurs := make(map[string]Currency)

	str, _ := document.Search("//p[@class='catatan']");
	firstData := strings.Split(strings.Split(str[0].InnerHtml(), "<br>")[0], " ")
	lastUpdated := firstData[2] + " " + firstData[3] + " " + firstData[4] + " " + firstData[5] + " " + firstData[6]

	doc, _ :=document.Search("//table[@class='tbl-view']/tr")
	for i, tr := range doc{
		t := 0
		for td := tr.FirstChild(); td != nil; td = td.NextSibling(){
			teks := strings.TrimSpace(td.Content())
			if i > 0 && i <= 15 && len(teks) > 0{
				if t == 1{
					matauang = teks
				} else if t == 2{
					kursbeli = teks
				} else if t == 3{
					kursjual = teks
				}
				t += 1
				kurs[matauang] = Currency{Nilai:nilai,KursJual:kursjual,KursBeli:kursbeli}
			}
		}
	}

	response := Data{
		LastUpdated: lastUpdated,
		Currency: kurs,
	}

	return response
}
