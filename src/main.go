package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Hello World")
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Ooops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(res, "Hello '%s'", data)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":3000", nil)
}
