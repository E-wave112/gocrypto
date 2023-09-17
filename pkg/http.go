package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"log/slog"

	"github.com/vicanso/go-axios"
)

const EXCHANGE_BASE_URL = "https://api.coinbase.com"

func LoggerMethod(name string, key string, value string) {
	f, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(value + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer f.Close()
	logger := slog.Default()
	logger.Info(name, key, string(value))

}

type MyBody struct {
	Data any `json:"data"`
}

func RetrieveRates(currency string) any {
	path := "/v2/exchange-rates?currency="
	url := fmt.Sprintf("%s%s%s", EXCHANGE_BASE_URL, path, currency)
	resp, err := axios.Get(url)
	if err != nil {
		panic(err)
	}
	LoggerMethod("json", "response", string(resp.Data))
	var data MyBody
	json.Unmarshal(resp.Data, &data)
	return data

}

func main() {
	currency := "BTC"
	response := RetrieveRates(currency)
	fmt.Println(response)
}
