package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	g := new(errgroup.Group)

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		// Launch a goroutine to fetch the URL
		url := url
		g.Go(func() error {
			// Fetch the URL
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:  response:", resp)
				resp.Body.Close()
				return err
			}
			fmt.Println(resp.Status)
			//fmt.Println(resp.Body)
			return err
		})
	}
	// Wait for all HTTP gets to complete
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs")
	}
}
