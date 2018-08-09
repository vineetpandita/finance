package main

import (
"net/http"
"io/ioutil"
"fmt"
"strings"
"encoding/json"
)

type NiftyStructWrap struct {
	Wrapper [] NiftyStruct `json:"d"`
}

type NiftyStruct struct {

	IndexName string `json:"INDEX_NAME"`
	HisDate string `json:"HistoricalDate"`
	ClosePrice string `json:"CLOSE"`
}

type date struct {
	f string
	t string
}

func main() {

//	indices := []string {"NIFTY 50","NIFTY Bank","NIFTY FMCG","NIFTY IT","NIFTY Infra","NIFTY Realty","NIFTY Pharma","NIFTY Auto","NIFTY Commodities","NIFTY Metal"}
	indices := []string {"NIFTY 50","NIFTY Next 50","NIFTY Midcap 150","NIFTY Smallcap 250"}

	dates := [] date {  {"15-Jun-2018","15-Jun-2018"},
		{"15-Jun-2017","15-Jun-2017"},
		{"15-Jun-2016","15-Jun-2016"},
		{"15-Jun-2015","15-Jun-2015"},
		{"16-Jun-2014","16-Jun-2014"},
		{"15-Jun-2007","15-Jun-2007"}}



	for _,index := range indices {

		println("")

		for _,dt := range dates {

			callService("{'name':'"+index+"','startDate':'"+dt.f+"','endDate':'"+dt.t+"'}")
		}

	}

	// "{'name':'+NIFTY 50','startDate':'15-May-2018','endDate':'31-May-2018'}"

}

func callService(body string) {
	err, body := fetchPost(body)
	if err != nil {
		fmt.Println("Failure in fethching data from URL")
		panic(err)
	}
	if err != nil {
		fmt.Println("Failure in parsing fethced data ")
		panic(err)
	} else {
		var m NiftyStructWrap
		body = cleanupJSON(body)
		e := unmarshallJson(body, &m)
		//fmt.Printf("\n JSON Object %+v ", m)

		if e != nil {
			panic(e)
		} else {
			//fmt.Println("\n Data Length  ", len(m.Wrapper))

			for _, val := range m.Wrapper {

				println(val.ClosePrice, ",", val.HisDate, ",", val.IndexName)
			}

			//str, _ := json.Marshal(m.Wrapper)
			//fmt.Printf("%+v", string(str))

		}

	}
}

func  fetchPost(bdy string) (error, string){

	body := strings.NewReader(bdy)
	req, err := http.NewRequest("POST",
		"http://www.niftyindices.com/Backpage.aspx/getHistoricaldatatabletoString", body)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(" Body ", string(resBody))
	return err, string(resBody)

}



func unmarshallJson(ss string, v interface{}) error {

	b := [] byte (ss);
	return json.Unmarshal(b, v)
}

func cleanupJSON(text string) string  {
	text = strings.Replace(text,"\"[","[",-1)
	text = strings.Replace(text,"]\"","]",-1)
	text = strings.Replace(text,"\\","",-1)
	return text
}



