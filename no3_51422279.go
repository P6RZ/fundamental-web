package main

import "fmt"

func main() {
	fmt.Print("enter a number: ")
	var input float32
	fmt.Scanf("%f", &input) //inputkan npm
	output := input * 2
	fmt.Println(output)
}
