package main

import (
	"context"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
)

func main() {
	ctx := context.Background()

	client, err := elastic.NewClient()

	if err != nil {
		panic(err)
	}

	fmt.Println(ctx, client)

	// ping the elastic search server to get e.g. the version number
	info, code, err := client.Ping("https://localhost:9201").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elastic search returned with code %d and version %s\n", code, info.Version.Number)
}
