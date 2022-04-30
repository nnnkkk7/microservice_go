package handlers

import (
	"log"
	"microservice_go/product-api/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}


func NewProducts(l *log.Logger) *Products{
	return &Products{}
}

func (p *Products)ServeHTTP(rw http.ResponseWriter, r *http.Request){
   if r.Method == http.MethodGet {
    p.getProducts(rw,r)
	return
   }
   
   rw.WriteHeader(http.StatusMethodNotAllowed)
	
}

func(p *Products)getProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	
}