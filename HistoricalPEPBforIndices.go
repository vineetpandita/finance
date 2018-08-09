package main


import (
"net/http"
"io/ioutil"
"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

type date1 struct {
	f string
	t string
}


type Pepbholder struct {

	HisDate string
	Pe string
	Pb string
	Dy string

}



func main() {

	//"NIFTY%2050","NIFTY%20BANK","NIFTY%20FMCG","NIFTY%20IT","NIFTY%20INFRASTRUCTURE",
		//indices := []string {"NIFTY%20PHARMA","NIFTY%20AUTO","NIFTY%20COMMODITIES","NIFTY%20METAL","NIFTY%20REALTY"}
	indices := []string {"NIFTY%20ENERGY","NIFTY%2050","NIFTY%20NEXT%2050","NIFTY%20MIDCAP%20150","NIFTY%20SMALLCAP%20250"}

	dates := [] date1 {
		{"31-07-2018","31-07-2018"},
		{"15-06-2018","15-06-2018"},
		{"15-06-2017","15-06-2017"},
		{"15-06-2016","15-06-2016"},
		{"15-06-2015","15-06-2015"},
		{"16-06-2014","16-06-2014"},
		{"14-06-2013","14-06-2013"},
		{"15-06-2007","15-06-2007"}}

	for _,index := range indices {
		println("\n ",index)
		for _,dt := range dates {
			val := callPEPBService("indexName="+index+"&fromDate="+dt.f+"&toDate="+dt.t+"&yield4=all")
			println( val.HisDate, ",", val.Pe, ",", val.Pb, ",", val.Dy)
			}

	}
	// "?indexName=NIFTY 50&fromDate=15-06-2018&toDate=15-06-2018&yield4=all"

}

func callPEPBService(params string) (Pepbholder) {
	err, body := fetchViaGet(params)

	var dt = ""
	var pe = ""
	var pb = ""
	var dy = ""


	if err != nil {
		fmt.Println("Failure in fethching data from URL")
		panic(err)
	} else {
		 // fmt.Println(body)

			doc, err := htmlquery.Parse(strings.NewReader(body))
			if err != nil {
				fmt.Println("Failure in parsing response")
			}


			for _, n := range htmlquery.Find(doc, "//tbody/tr/td[@class='date']") {
				dt = htmlquery.InnerText(n)
			}
			for i, n := range htmlquery.Find(doc, "//tbody/tr/td[@class='number']") {

				switch i {

				case 0 : pe = htmlquery.InnerText(n) ; break
				case 1 : pb = htmlquery.InnerText(n) ; break
				case 2 : dy = htmlquery.InnerText(n)

				}
			}
		}


		vo  :=Pepbholder{dt,pe,pb,dy}


		return vo

	}


func  fetchViaGet(params string) (error, string){

	url := "https://www.nseindia.com/products/dynaContent/equities/indices/historical_pepb.jsp?"+params

	//fmt.Println(url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)

	if err != nil {
		return err, ""
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(" Body ", string(resBody))
	return err, string(resBody)

}








