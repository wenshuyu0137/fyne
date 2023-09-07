package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var currency = "USD"

type Price struct {
	Currency       string    `json:"currency"`
	Price          float64   `json:"xauPrice"`
	Change         float64   `json:"chgXau"`
	Previous_close float64   `json:"xauClose"`
	Time           time.Time `json:"-"`
}

type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

func (g *Gold) get_prices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client

	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)

	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error contacting goldprice.org", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading json", err)
		return nil, err
	}

	gold := Gold{}

	var previous, current, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("error ummarshal", err)
		return nil, err
	}

	previous, current, change = gold.Prices[0].Previous_close, gold.Prices[0].Price, gold.Prices[0].Change

	var current_info = Price{
		Currency:       currency,
		Price:          current,
		Change:         change,
		Previous_close: previous,
		Time:           time.Now(),
	}

	return &current_info, nil

}
