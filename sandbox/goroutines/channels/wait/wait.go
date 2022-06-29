package wait

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func WelcomeMessage() {
	fmt.Println("Welcome to Educative!")
}

func WaitMain() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		WelcomeMessage()
		wg.Done()
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		fmt.Println("Hello world")
		wg.Done()
	}()

	wg.Wait()
}
