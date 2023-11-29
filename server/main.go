package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/E-wave112/gocrypto/pkg"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
)

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	val, ok := pkg.UnsetValueInRedis("safin@outlook.com")
	// redisVal := "up and running!"

	io.WriteString(w, fmt.Sprintf("%s : %v\n", val, ok))

}

func main() {

	viper.SetConfigFile(".env")

	// Find and read the config file
	viperErr := viper.ReadInConfig()

	if viperErr != nil {
		log.Fatalf("Error while reading config file %s", viperErr)
	}

	port, ok := viper.Get("PORT").(string)

	if !ok {
		port = ":9000"
	}
	scheduler := gocron.NewScheduler(time.UTC)
	_, jobErr := scheduler.Every("30m").Do(func() {
		pkg.LoggerMethod("cronservice", "cron", "exchange rate background service is starting....")
		currency := "USD"
		rates, _ := pkg.RetrieveRates(currency)
		fmt.Println(rates)
		// TODO: function to retrieve and send emails to subscribers
	})

	if jobErr != nil {
		newErr := errors.New("an error occurred when starting the cron service")
		fmt.Println(newErr)
	}

	scheduler.StartAsync()
	mux := http.NewServeMux()
	mux.HandleFunc("/check", getHealthCheck)

	log.Printf("server listening on port %s\n", port)
	err := http.ListenAndServe(port, mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		os.Exit(1)
	}
}
