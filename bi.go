package kurs

import(
	"strings"
	"github.com/moovweb/gokogiri/html"
)

func processBI() Data{
	url := "http://www.bi.go.id/id/moneter/informasi-kurs/transaksi-bi/Default.aspx"
	
	document := processUrl(url)
	defer document.Free();

	lastUpdated := getLastUpdatedBi(document)

 	data,_ := parseBiHtml(document)

	response := Data{
		LastUpdated: lastUpdated,
		Currency: data,
	}

	return response
}

func getLastUpdatedBi(document *html.HtmlDocument) string {
	span, _ := document.Search("//span[@id='ctl00_PlaceHolderMain_biWebKursTransaksiBI_lblUpdate']")
	lastUpdated := span[0].InnerHtml()

	return lastUpdated
}

func parseBiHtml(document *html.HtmlDocument) (map[string]Currency, error) {
	var matauang string
	var nilai string
	var kursjual string
	var kursbeli string
	
	kurs := make(map[string]Currency)

	doc, err := document.Search("//table[@id='ctl00_PlaceHolderMain_biWebKursTransaksiBI_GridView1']/tr")
	for i, tr := range doc{
		t := 0
		for td := tr.FirstChild(); td != nil; td = td.NextSibling(){
			teks := strings.TrimSpace(td.Content())
			if i > 0 && len(teks) > 0{
				if t == 0{
					matauang = teks
				} else if t == 1{
					nilai = teks
				} else if t == 2{
					kursjual = teks
				} else if t == 3{
					kursbeli = teks
				}
				t += 1
				kurs[matauang] = Currency{Nilai:nilai,KursJual:kursjual,KursBeli:kursbeli}
			}
		}
	}

	return kurs, err
}