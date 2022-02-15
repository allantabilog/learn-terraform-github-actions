package main

import (
	"fmt"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	ctx = addValue(ctx)
	readValue(ctx)

}

func addValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "key", "test value")
}

func readValue(ctx context.Context) {
	val := ctx.Value("key")
	fmt.Println(val)
}
