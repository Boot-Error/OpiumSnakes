package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {

	fmt.Println("snekonplane")

	http.HandleFunc("/start", StartHandler)
	http.HandleFunc("/end", EndHandler)
	http.HandleFunc("/move", MoveHandler)
	http.HandleFunc("/ping", PingHandler)

	port := os.Getenv("PORT")

	fmt.Println(
		http.ListenAndServe(strings.Join([]string{"0.0.0.0", port}, ":"), nil))
}
