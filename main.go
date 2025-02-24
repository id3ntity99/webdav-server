package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    http.Handle("/", http.HandlerFunc(func (w http.ResponseWriter, r * http.Request) {
        urlStr := r.URL.String();
        fmt.Println("User hit " + urlStr);
    }))

    log.Fatal(http.ListenAndServe(":8001", nil))
}