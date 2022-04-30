package main

import (
	"context"
	"log"
	"microservice_go/product-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout,"product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil  {
			l.Fatal(err)
		}
		os.Exit(1)
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
