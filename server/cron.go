package main

import "github.com/E-wave112/gocrypto/pkg"

func schedulerForCryptoExchangeRates() pkg.Rates {
	currency := "USD"
	rates, _ := pkg.RetrieveRates(currency)
	// if err != nil {
	// 	return "could not retrieve the rates at this time"
	// }
	return rates
}
