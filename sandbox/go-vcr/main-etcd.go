package main

import (
	"github.com/dnaeon/go-vcr/v2/recorder"
	"go.etcd.io/etcd/client/v2"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {
	// Start the recorder
	r, err := recorder.New("fixtures/etcd")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Stop() // Ensure that the recorder is stopped once we're done with it

	// Create an etcd configuration using our transport
	cfg := client.Config {
		Endpoints:	[]string{"http://127.0.0.1:2379"},
		HeaderTimeoutPerRequest: time.Second,
		Transport: r,
	}
	
	c, err := client.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create etcd client: %s", err)
	}
	
	// Get an example key from etcd
	etcdKey := "/foo"
	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Get(context.Background(), etcdKey, nil)
	if err != nil {
		log.Fatalf("Failed to create etcd key %s: %s", etcdKey, err)
	}
	log.Printf("Successfully retrieved etcd key: %s %s", etcdKey, resp.Node.Value)
	
	
}
