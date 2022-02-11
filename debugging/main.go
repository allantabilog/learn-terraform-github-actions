package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	readTimeout  = 5
	writeTimeout = 10
	idleTimeout  = 120
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	w.WriteHeader(returnStatus)
	message := fmt.Sprintf("Hello %s!", r.UserAgent())
	w.Write([]byte(message))
}

func main(){
	serverAddress :=  ":8080"
	l := log.New(os.Stdout, "sample-srv", log.LstdFlags | log.Lshortfile)
	m := mux.NewRouter()

	m.HandleFunc("/", indexHandler)

	srv :=  &http.Server{
		Addr: serverAddress,
		ReadTimeout: readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout: idleTimeout,
		Handler: m,
	}

	l.Println("server started")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}