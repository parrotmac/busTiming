package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const MSG_HTTP_FAILURE = "BAD REQ"
const MSG_PARSE_FAILURE = "PARSE ERR"
const MSG_NO_ROUTES = "0 ROUTES"

type DisplayMessageDef struct {
	Messages	[]string

}

type BusThingy struct {
	StopID	string
	Agency	string
}

type DirectionalPrediction struct {
	AgencyTitle	string	`xml:"agencyTitle,attr"`
	Title	string `xml:"title,attr"`
	Predictions	[]struct{
		EpochTime	int64	`xml:"epochTime,attr"`
		Minutes		int64	`xml:"minutes,attr"`
		Vehicle		int64	`xml:"vehicle,attr"`
	}	`xml:"prediction"`
}

type TransitPredictionResult struct {
	XMLName	xml.Name `xml:"body"`
	Copyright	string	`xml:"copyright,attr"`
	Predictions []struct {
		DirectionPredictions []DirectionalPrediction `xml:"direction"`
	} `xml:"predictions"`
}

func (b* BusThingy) getTargetURL() string {
	//myStopURL := "http://webservices.nextbus.com/service/publicXMLFeed?command=predictions&a=actransit&stopId=58758"
	return fmt.Sprintf(
		"http://webservices.nextbus.com/service/publicXMLFeed?command=predictions&a=%s&stopId=%s",
		b.Agency,
		b.StopID,
	)
}

func (b* BusThingy) getNextBusResponseBody(apiURL string) ([]byte, error, string) {
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		return []byte(""), err, MSG_HTTP_FAILURE
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return []byte(""), err, MSG_PARSE_FAILURE
	}
	return body, nil, ""
}

func (b* BusThingy) getDepartureEpochMinutes() ([]int, error) {
	return [0, 1], nil
}

func main() {

	busCfg := &BusThingy{
		StopID: "58758",
		Agency: "actransit",
	}

	apiURL := busCfg.getTargetURL()
	timingXML, err, shortErrMessage := busCfg.getNextBusResponseBody(apiURL)

//	exampleData := `<?xml version="1.0" encoding="utf-8" ?>
//<body copyright="All data copyright AC Transit 2018.">
//<predictions agencyTitle="AC Transit" routeTitle="C" routeTag="C" stopTitle="41st St &amp; Piedmont Av" stopTag="1002710">
//<direction title="To San Francisco">
//<prediction epochTime="1537284708203" seconds="874" minutes="14" isDeparture="false" affectedByLayover="true" dirTag="C_27_1" vehicle="6055" block="103004" tripTag="6242853" />
//</direction>
//</predictions>
//</body>`


	ed := TransitPredictionResult{}

	err = xml.Unmarshal([]byte(timingXML), &ed)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(ed)
	
}
