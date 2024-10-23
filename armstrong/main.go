package main

import (
	"fmt"
	"math"
)

func main() {

	tests := []int{153, 370, 371, 407, 9474, 9475}

	for _, test := range tests {
		fmt.Printf("The number %v is %v\n", test, isArmstrongV2(test))
	}
}

func isArmstrong(number int) string {

	// check if the number is a 3 digit number
	if number < 100 || number > 999 {
		return "InputError"
	}

	// split the number into digits
	var digits []int
	digits = splitIntoDigits(number)

	// fmt.Println("The digits are : %v", digits)

	// sum the cube of each digit
	var sum float64
	for _, digit := range digits {
		sum += math.Pow(float64(digit), 3)
	}

	// fmt.Printf("The sum is : %v", sum)
	// check if the sum equals the original number
	if int(sum) == number {
		return "Armstrong"
	} else {
		return "NotArmstrong"
	}
}

func splitIntoDigits(number int) []int {
	var digits []int
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}
	return digits
}

func isArmstrongV2(number int) string {
	if number > 999 || number < 100 {
		return "InputError"
	}

	// get the digits of the number
	unit := number % 10
	tens := (number / 10) % 10
	hundreds := (number / 100) % 10

	// check if the number is an armstrong number
	sumCubes := math.Pow(float64(unit), 3) +
		math.Pow(float64(tens), 3) +
		math.Pow(float64(hundreds), 3)

	if int(sumCubes) == number {
		return "Armstrong"
	}

	return "NotArmstrong"

}
