package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Bad request data", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s.\n", b)
	})

	http.HandleFunc("/healthCheck", func(rw http.ResponseWriter, _ *http.Request) {
		log.Println("UP!!!")
		fmt.Fprintf(rw, "UP! \n")
	})

	http.ListenAndServe(":9090", nil)
}
