package main

import "fmt"

//program execution starts here
func main(){

  //declare a map
  students := map[string]int{
    "john": 10,
    "james": 12,
    "mary":13,
  }
  
  //provide user value 
  var userValue = 10

//   
_, ok := students[userValue]
if(ok) {
	fmt.Println("Value found")
} else {
	fmt.Println("Value not found")
}

  
}

//function to check if a value present in the map
func checkForValue(userValue int, students map[string]int)bool{

  //traverse through the map
  for _,value:= range students{

    //check if present value is equals to userValue
    if(value == userValue){

      //if same return true
      return true
    }
  }

  //if value not found return false
  return false
}
