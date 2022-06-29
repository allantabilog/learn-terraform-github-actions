package channel

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func ContextMain() {
	http.ListenAndServe(":8001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		fmt.Fprintf(os.Stdout, "processing request\n")

		select {
		case <-time.After(time.Second * 60):
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled
			fmt.Fprintf(os.Stderr, "request cancelled")
		}
	}))
}
