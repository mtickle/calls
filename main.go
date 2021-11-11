package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.86.2"
	port     = 5432
	user     = "pi"
	password = "Boomer2025"
	dbname   = "incident"
)

func OnPage(link string) {

	//-----------------------------------------------------------------------
	//--- Scrape the page into a value locally.
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Make and open the database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//--- Are we good?
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Grab the JSON sites data and split out the events
	var temp_sites = GetStringInBetween(string(content), " var sites = [", "];")
	sites := strings.Split(temp_sites, "],")
	//-----------------------------------------------------------------------

	//-----------------------------------------------------------------------
	//--- Loop through the array. This is where the magic happens.
	for _, s := range sites {
		var temp_site = ""

		//--- Split once at the commas
		temp_site = strings.Replace(s, "[", "", -1)
		site := strings.Split(temp_site, ",")

		//--- Work up the longer string of HTML at the end and split it.
		var temp_location = site[4]
		//fmt.Println(temp_location)
		temp_location = strings.Replace(temp_location, "<br /><br />Loc: ", "|", -1)
		temp_location = strings.Replace(temp_location, "<br />Time: ", "|", -1)
		temp_location = strings.Replace(temp_location, "<br />Agency: ", "|", -1)
		temp_location = strings.Replace(temp_location, "'<strong>", "", -1)
		temp_location = strings.Replace(temp_location, "</strong>", "", -1)
		call := strings.Split(temp_location, "|")

		//--- Clarify everything.
		var Agency = strings.Replace(call[3], "'", "", -1)
		var Latitude = site[1]
		var Longitude = site[2]
		var Incident = call[0]
		var Location = call[1]
		var Datestamp = call[2]

		//--- Run your insert.
		var sql = "CALL add_call ($1, $2, $3, $4, $5, $6);"
		_, err := db.Exec(sql, Agency, Latitude, Longitude, Incident, Location, Datestamp)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(Agency)
	}
	//-----------------------------------------------------------------------

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
