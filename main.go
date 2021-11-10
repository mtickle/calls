package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Site struct {
	Latitude  string
	Longitude string
	Agency    string
	Address   string
	Incident  string
	Datestamp string
}

var arraySite []Site

func OnPage(link string) {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var temp_sites = GetStringInBetween(string(content), " var sites = [", "];")

	sites := strings.Split(temp_sites, "],")
	arrSite := Site{}
	//arrSite := Site{}

	//--- Loop through the array. This is where the magic happens.
	for _, s := range sites {
		var temp_site = ""
		temp_site = strings.Replace(s, "[", "", -1)
		site := strings.Split(temp_site, ",")

		arrSite.Agency = site[0]
		arrSite.Latitude = site[1]
		arrSite.Longitude = site[2]

		var temp_location = site[4]
		temp_location = strings.Replace(temp_location, "<br /><br />Loc: ", "|", -1)
		temp_location = strings.Replace(temp_location, "<br />Time: ", "|", -1)
		temp_location = strings.Replace(temp_location, "<br />Agency: ", "|", -1)
		temp_location = strings.Replace(temp_location, "<strong>", "", -1)
		temp_location = strings.Replace(temp_location, "</strong>", "", -1)

		call := strings.Split(temp_location, "|")

		arrSite.Incident = call[0]
		arrSite.Address = call[1]
		arrSite.Datestamp = call[2]

		//temp_location = strings.Replace(temp_location, "<br /><br />", "|", -1)
		//temp_location = strings.Replace(temp_location, "<br />", "|", -1)
		//arrSite.Address

		for i, s := range arrSite {
			fmt.Println(i, s)
		}

		//fmt.Println(arrSite)
	}

}

func main() {
	OnPage("http://incidents.rwecc.com/")

}

func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}

// func ExampleScrape() {
// 	// Request the HTML page.
// 	doc, err := goquery.NewDocument("http://incidents.rwecc.com/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	doc.Find("table").Each(func(t int, tr *goquery.Selection) {

// 		switch t {
// 		case 0:
// 		case 1:
// 			s := Site{}

// 			tr.Find("td").Each(func(ix int, td *goquery.Selection) {

// 				//fmt.Println(strconv.Itoa(ix) + "  " + td.Text())

// 				switch ix {
// 				case 1:
// 					s.Agency = td.Text()
// 				case 2:
// 					s.Address = td.Text()
// 				case 3:
// 					s.Incident = td.Text()
// 				case 4:
// 					s.Datestamp = td.Text()
// 				}
// 				//fmt.Println(s.Agency)
// 				arraySite = append(arraySite, s)
// 			})
// 			//fmt.Println(s.Agency)
// 			//arraySite = append(arraySite, s)
// 		}

// 	})

// 	//fmt.Println(arraySite)

// 	for i, s := range arraySite {
// 		fmt.Println(i, s)
// 	}

// }

// func main() {
// 	ExampleScrape()
// }
