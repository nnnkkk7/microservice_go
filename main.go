package main

import (
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request){
		log.Println("hello world")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("data %s\n", d)
	})
	http.ListenAndServe(":9090", nil)
}