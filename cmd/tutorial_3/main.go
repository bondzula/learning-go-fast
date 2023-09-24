package main

import "fmt"

func main() {
	printMe("Hello world")

	numerator := 11
	denominator := 2
	var result, remainder int = intDivision(numerator, denominator)

	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}
