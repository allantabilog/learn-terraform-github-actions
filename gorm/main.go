package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	newProduct := Product{Code: "D42", Price: 100}
	result := db.Create(&newProduct)

	fmt.Printf("Product created with ID: %d", newProduct.ID)
	fmt.Printf("Result returned: %#v", result)

	newProduct2 := Product{Code: "D43", Price: 150}
	result = db.Select("Code", "Price").Create(&newProduct2)

	fmt.Printf("Product created with ID: %d", newProduct2.ID)
	fmt.Printf("Result returned: %#v", result)

	// Read
	// var product Product
	// fmt.Println(db.First(&product, 1))
	// fmt.Println(db.First(&product, "code = ?", "D42"))

	// //Update
	// db.Model(&product).Update("Price", 200)
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

}
