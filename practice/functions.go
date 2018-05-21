package main

import "fmt"


func myName() string {
	return "Robin Smith"
}

func myAge() int {
	return 30
}

func main() {
	fmt.Println("I am " + myName() + " and I am " + fmt.Sprintf("%d", myAge()) + " years old.")
}