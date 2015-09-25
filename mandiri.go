package kurs

import(
	"strings"
	"github.com/moovweb/gokogiri/html"
)

func processMandiri() Data{
	url := "http://www.bankmandiri.co.id/resource/kurs.asp?row=2"
	
	document := processUrl(url)
	defer document.Free();

	lastUpdated := getLastUpdated(document)

	data,_ := parseMandiriHtml(document)

	response := Data{
		LastUpdated: lastUpdated,
		Currency: data,
	}

	return response
}

func getLastUpdated(document *html.HtmlDocument) string {
	str, _ := document.Search("//p[@class='catatan']");
	firstData := strings.Split(strings.Split(str[0].InnerHtml(), "<br>")[0], " ")
	lastUpdated := firstData[2] + " " + firstData[3] + " " + firstData[4] + " " + firstData[5] + " " + firstData[6]

	return lastUpdated
}

func parseMandiriHtml(document *html.HtmlDocument) (map[string]Currency, error) {
	nilai := "1.00"
	var matauang string
	var kursjual string
	var kursbeli string

	var kurs = make(map[string]Currency)
	doc, err :=document.Search("//table[@class='tbl-view']/tr")
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
 
				if kursjual != "" && kursbeli != "" {
					kurs[matauang] = Currency{Nilai:nilai,KursJual:kursjual,KursBeli:kursbeli}
				}
			}
		}
	}

	return kurs, err
}
