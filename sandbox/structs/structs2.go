package main

import "fmt"

type Record struct {
	Name string
	Age  int
}

func (r Record) String() string {
	return fmt.Sprintf("%s is %d years old", r.Name, r.Age)
}

func NewRecord(name string, age int) (*Record, error) {
	if age < 0 {
		return nil, fmt.Errorf("age cannot be negative")
	}
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	return &Record{Name: name, Age: age}, nil
}

func main() {
	record := Record{
		Name: "John",
		Age:  25,
	}

	fmt.Println(record.String())

	rec1, err := NewRecord("Jane", 30)
	if err != nil {
		fmt.Println(err)
	}
	rec2, err := NewRecord("John", 31)
	if err != nil {
		fmt.Println(err)
	}

	records := []*Record{rec1, rec2}

	for _, record := range records {
		fmt.Println(record.String())
	}
}
