package main

import "fmt"


func main() {
	const a = 1
	fmt.Println("a: " + fmt.Sprintf("%d", a))
	var b = 2
	fmt.Println("b: " + fmt.Sprintf("%d", b))
	b += 1
	fmt.Println("b: " + fmt.Sprintf("%d", b))
	const c = 1.23
	fmt.Println("c: " + fmt.Sprintf("%f", c))

	const d = 1.23

	const e = c == d
	fmt.Println("e: " + fmt.Sprintf("%t", e))
}