package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://carbonfootprint1.p.rapidapi.com/CarbonFootprintFromCarTravel?distance=100&vehicle=SmallDieselCar"

	//url := "https://google.com"
	//url := "https://lco.dev"

	req, err := http.NewRequest("GET", url, nil)
	//req, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-RapidAPI-Key", "8f372d52a8msh5b9395d483bced2p13d066jsn3cba779ae35b")
	req.Header.Add("X-RapidAPI-Host", "carbonfootprint1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	//fmt.Println(res)
	fmt.Println(string(body))

}
