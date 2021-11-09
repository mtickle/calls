package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Site struct {
	Agency    string
	Address   string
	Incident  string
	Datestamp string
}

var arraySite []Site

func ExampleScrape() {
	// Request the HTML page.
	doc, err := goquery.NewDocument("http://incidents.rwecc.com/")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table").Each(func(t int, tr *goquery.Selection) {
		switch t {
		case 0:
		case 1:
			s := Site{}

			tr.Find("td").Each(func(ix int, td *goquery.Selection) {
				fmt.Println(ix)
				switch ix {
				case 1:
					s.Agency = td.Text()
				case 2:
					s.Address = td.Text()
				case 3:
					s.Incident = td.Text()
				case 4:
					s.Datestamp = td.Text()
				}
			})

			arraySite = append(arraySite, s)
		}

	})

	fmt.Println(arraySite)

}

func main() {
	ExampleScrape()
}
