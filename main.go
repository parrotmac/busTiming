package main

import (
	"encoding/xml"
	"fmt"
)

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

func main() {
	_ = "http://webservices.nextbus.com/service/publicXMLFeed?command=predictions&a=actransit&stopId=58758"
	_ = ""

	exampleData := `<?xml version="1.0" encoding="utf-8" ?>
<body copyright="All data copyright AC Transit 2018.">
<predictions agencyTitle="AC Transit" routeTitle="C" routeTag="C" stopTitle="41st St &amp; Piedmont Av" stopTag="1002710">
<direction title="To San Francisco">
<prediction epochTime="1537284708203" seconds="874" minutes="14" isDeparture="false" affectedByLayover="true" dirTag="C_27_1" vehicle="6055" block="103004" tripTag="6242853" />
</direction>
</predictions>
</body>`

	ed := TransitPredictionResult{}

	err := xml.Unmarshal([]byte(exampleData), &ed)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(ed)
	
}
