package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/E-wave112/gocrypto/pkg"
	"github.com/robfig/cron"
)

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "the system is up and running!\n")

}

func main() {

	c := cron.New()

	c.AddFunc("", func() {
		currency := "USD"
		rates, _ := pkg.RetrieveRates(currency)
		// if err != nil {
		// 	return "could not retrieve the rates at this time"
		// }
		fmt.Println(rates)
	})
	mux := http.NewServeMux()
	mux.HandleFunc("/check", getHealthCheck)

	// Start the Cron job scheduler
	c.Start()

	err := http.ListenAndServe(":9000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		os.Exit(1)
	}
}
