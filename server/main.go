package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "the system is up and running!\n")

}

func main() {
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
