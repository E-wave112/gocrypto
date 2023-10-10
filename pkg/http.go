package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"strconv"

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

// The LoggerMethod function appends a value to a log file named "info.log".
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
	// var logger = slog.Default()
	// // logger.Info(name, key, string(value))

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

type Response struct {
	Data RateRep `json:"data"`
}

type Rates map[string]string

// The `formatNumber` function takes a string input, converts it to a float, calculates the inverse,
// formats it as a significant figure, and returns the formatted output.
func formatNumber(input string) string {
	floatInput, _ := strconv.ParseFloat(input, 64)
	// get the inverse of the floatInput
	floatInput = 1 / floatInput
	floatToString := fmt.Sprintf("%17f", floatInput)
	formattedOutput := strings.TrimRight(floatToString, "0")
	// Remove trailing "." if present
	if strings.HasSuffix(formattedOutput, ".") {
		formattedOutput = formattedOutput[:len(formattedOutput)-1]
	}
	// return the `formattedoutput` as a significant figure
	return formattedOutput
}

// The function `RetrieveRates` retrieves the exchange rate between the supported cryptocurrencies and any specified fiat currency provided as an argument and
// returns the results in a structured format.
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
		"Bitcoin":  formatNumber(rates.Rates.BTC),
		"Ethereum": formatNumber(rates.Rates.ETH),
		"Dogecoin": formatNumber(rates.Rates.DOGE),
		"Solana":   formatNumber(rates.Rates.SOL),
		"Shiba":    formatNumber(rates.Rates.SHIB),
		"Tether":   formatNumber(rates.Rates.USDT),
	}
	return context, nil

}

// lists the supported cryptocurrencies to get real-time exchange rates from
func ListSupportedCryptoCurrencies() []string {
	supportedCurrencies := []string{"Bitcoin (BTC)", "Ethereum(ETH)", "Dogecoin(DOGE)", "Solana(SOL)", "Shiba(SHIB)", "Tether(USDT)"}
	return supportedCurrencies
}

func main() {
	currency := "USD"
	response, _ := RetrieveRates(currency)
	fmt.Println(response)
}
