package main

// sample code and article
// from: https://blog.logrocket.com/rate-limiting-go-application/
// also check out:
// Tollbooth - go package for rate limiting (no need to roll your own)
// https://github.com/didip/tollbooth
import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/time/rate"
)

type Message struct {
	Status string `json:"status"`	
	Body 	string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, req *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	message := Message {
		Status: "Successful",
		Body: "HELLO from the API",
	}

	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}

}

// just an http middleware
func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request ) {
		if !limiter.Allow() {
			message := Message {
				Status: "Request Failed",
				Body: "The API is at capacity, try again later.",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		} else {
			next(w, r)
		}
	})

}

func main() {
	http.Handle("/ping", rateLimiter(endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error listening on port 8080", err)
	}
}