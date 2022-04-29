package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *GoodBye{
 return &GoodBye{}
}

func(g *GoodBye)ServeHTTP(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("bye"))
}