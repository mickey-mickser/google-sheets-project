package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting bot...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, bot!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
