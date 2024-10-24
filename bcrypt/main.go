package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func getPwd() []byte {

	fmt.Println("Enter password:")
	var pwd string 

	_, err := fmt.Scan(&pwd)

	if err != nil {
		log.Println(err)
	}

	return []byte(pwd)
	
}

func hashAndSalt(pwd []byte) string {
	
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool{

	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		log.Println(err)
		return false
	}

	return true



}

func main(){

	for {
		pwd := getPwd()
		hash := hashAndSalt(pwd)
		pwd2 := getPwd()

		pwdMatch := comparePasswords(hash, pwd2)

		fmt.Println("Passwords match?", pwdMatch)

	}
	
}