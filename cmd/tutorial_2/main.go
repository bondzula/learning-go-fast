package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// int int8 int16 int32 int64 uint uint8...
	var intNum int = 32767
	fmt.Println(intNum)

	// float32 and float64, where float64 is more precise
	var floatNum float64
	fmt.Println(floatNum)

	// Aritmetic operations can couse overflow, is our initial int is to small for example

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// To add those 2 types, we need to convert the int first
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // 1 rounded down
	fmt.Println(intNum1 % intNum2) // 1 remainder

	var myString string = "Hello" + " " + "World"
	// with back quotes, we can format the string

	fmt.Println(myString)
	fmt.Println(len(myString)) // this returns the number of bytes, not the number of characters

	fmt.Println(utf8.RuneCountInString(myString)) // Actualy returns number of characters

	// Rune is a data type in go, and represents a character
	var myRune rune = 'a'
	fmt.Println(myRune)

	// booleans are like always
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// In case the value is not set, go sets the default one
	// For numbers its 0
	// For strings is ''
	// And for booleans is false

	// We dont have to set the type of a variable if we set its value, and go can infer
	var myVar = "text"
	fmt.Println(myVar)

	// Shorthand syntax, without var
	myFancyVar := "text"
	fmt.Println(myFancyVar)

	// set multiple variables at onece
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// Same as variables, but can't be changed
	// And value has to be set
	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
}
