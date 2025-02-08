package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var logger TransactionLogger

func main() {
	if err := initializeTransactionLogger(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/v1/key/{key}", kvGetHandler).Methods("GET")
	r.HandleFunc("/v1/key/{key}", kvPutHandler).Methods("PUT")
	r.HandleFunc("/v1/key/{key}", kvDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func initializeTransactionLogger() error {
	var err error

	logger, err := NewFileTransactionLogger("transaction.log")
	if err != nil {
		return fmt.Errorf("failed to initialize transaction logger: %v", err)
	}

	events, errors := logger.ReadEvents()

	e := Event{}
	ok := true

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case EventPut:
				err = Put(e.Key, e.Value)
			case EventDelete:
				err = Delete(e.Key)
			}
		}
	}

	logger.Run()

	return err
}
