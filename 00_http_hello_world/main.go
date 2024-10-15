package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World!. You've requested: %s\n", request.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
