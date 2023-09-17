package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"log/slog"

	"github.com/vicanso/go-axios"
)

const (
	ErrBadRequest = Err("An error occured when trying to fetch exchange rates")
)

type Err string

func (e Err) Error() string {
	return string(e)
}

const EXCHANGE_BASE_URL = "https://api.coinbase.com"

func LoggerMethod(name string, key string, value string) {
	f, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(value + "\n\n\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer f.Close()
	logger := slog.Default()
	logger.Info(name, key, string(value))

}

type Currency struct {
	BTC  string
	ETH  string
	DOGE string
	SOL  string
	SHIB string
	USDT string
}
type RateRep struct {
	Rates Currency `json:"rates"`
}

// type Data struct {
// 	Currency string            `json:"currency"`
// 	Rates    map[string]string `json:"rates"`
// }

// // Define a struct to represent the entire JSON response
// type Response struct {
// 	Data Data `json:"data"`
// }

type Response struct {
	Data RateRep `json:"data"`
}

type Rates map[string]string

func RetrieveRates(currency string) (Rates, error) {
	path := "/v2/exchange-rates?currency="
	url := fmt.Sprintf("%s%s%s", EXCHANGE_BASE_URL, path, currency)
	resp, err := axios.Get(url)
	if err != nil {
		return Rates{}, ErrBadRequest
	}
	LoggerMethod("json", "response", string(resp.Data))
	var data Response
	json.Unmarshal(resp.Data, &data)
	rates := data.Data
	context := Rates{
		"Bitcoin":  rates.Rates.BTC,
		"Ethereum": rates.Rates.ETH,
		"Dogecoin": rates.Rates.DOGE,
		"Solana":   rates.Rates.SOL,
		"Shiba":    rates.Rates.SHIB,
		"Tether":   rates.Rates.USDT,
	}
	return context, nil

}

func main() {
	currency := "USD"
	response, _ := RetrieveRates(currency)
	fmt.Println(response)
}
