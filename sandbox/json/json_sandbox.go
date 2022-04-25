package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func sandboxMain() {
	//address()
	name()
}

type User struct {
	Name          string    `json:"name"`
	Password      string    `json:"-"`
	PreferredFish []string  `json:"preferredFish,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (u *User) String() string {
	return fmt.Sprintf("Hi my name is %s", u.Name)
}

func name() {
	u := User{Name: "Joe", Password: "joedimaggio123",
		CreatedAt: time.Now()}
	fmt.Println(u)

	out, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}

func address() {
	var addressJsonBlob = []byte(`[
		{"StreetNumber": 1, "Street": "Sesame", "StreetType": "Street", "Suburb": "Brooklyn", "Country": "USA"},
		{"StreetNumber": 33, "Street": "City", "StreetType": "Road", "Suburb": "Southbank", "Country": "Australia"}
]`)

	type Address struct {
		StreetNumber int
		StreetName   string
		StreetType   string
		Suburb       string
		Country      string
	}
	var addresses []Address
	err := json.Unmarshal(addressJsonBlob, &addresses)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%+v", addresses)
	fmt.Println(addresses[0].Country)

	nzAddr := Address{StreetNumber: 8, StreetName: "Chapman", StreetType: "St", Suburb: "Johnsonville", Country: "New Zealand"}
	jsonBytes, err := json.Marshal(nzAddr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	os.Stdout.Write(jsonBytes)

}
