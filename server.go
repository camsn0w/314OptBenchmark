package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Trip struct {
	RequestType    string `json:"requestType"`
	RequestVersion int    `json:"requestVersion"`
	Options        struct {
		Title       string `json:"title"`
		EarthRadius string `json:"earthRadius"`
		Response    string `json:"response"`
	} `json:"options"`
	Places []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Municipality string `json:"municipality"`
		Country      string `json:"country"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		Altitude     string `json:"altitude"`
	} `json:"places"`
}

func readToStruct(filename string) (Trip, error) {
	var result Trip
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(file, &result)
	return result, err
}

func genCases(trip Trip) []Trip {
	tripArr := make([]Trip, 5)
	for i := range tripArr {
		tripArr[i] = trip
	}
	for i := 0; i < 5; i++ {
		tripArr[i].Options.Response = strconv.FormatInt(int64(i)*2, 10)
	}
	return tripArr
}

func runCases(n int, filename string) []int64 {
	startTrip, err := readToStruct(filename)
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	cases := genCases(startTrip)
	times := make([]int64, len(cases))
	for i := range cases {
		start := time.Now()
		for j := 0; j < n; j++ {
			postJson(cases[i])
		}
		times[i] = time.Since(start).Milliseconds() / int64(n)

	}
	return times
}

func structToString(trip Trip) (string, error) {
	out, err := json.Marshal(trip)
	if err != nil {
		return "", err
	}

	return string(out), err
}

func postJson(trip Trip) int8 {
	url := "http://localhost:8000/api/trip?=&="
	tripStr, err := structToString(trip)
	tripReader := strings.NewReader(tripStr)
	if err != nil {
		print(err.Error())
	}

	req, _ := http.NewRequest("POST", url, tripReader)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	if res.Status != "HTTP/1.1 200 OK" {
		println(res.Status)
		os.Exit(-1)
	}

	return 0

}
