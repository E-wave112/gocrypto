package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/E-wave112/gocrypto/pkg"
	"github.com/go-co-op/gocron"
)

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "the system is up and running!\n")

}

func main() {
	scheduler := gocron.NewScheduler(time.UTC)
	_, jobErr := scheduler.Every("2m").Do(func() {
		pkg.LoggerMethod("cronservice", "cron", "cc rate background service is starting....")
		currency := "USD"
		rates, _ := pkg.RetrieveRates(currency)
		fmt.Println(rates)
	})

	if jobErr != nil {
		newErr := errors.New("an error occurred when starting the cron service")
		fmt.Println(newErr)
	}

	scheduler.StartAsync()
	mux := http.NewServeMux()
	mux.HandleFunc("/check", getHealthCheck)

	err := http.ListenAndServe(":9000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		os.Exit(1)
	}
}
