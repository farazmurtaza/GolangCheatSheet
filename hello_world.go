package main

import "fmt"

/* this is a multiline
comment */

func main() {
	// fmt.Println("Hello world")

	// variable name dataType
	var age int = 21

	fmt.Println(age)

	var pi float64 = 3.1415923

	// fmt.Println(pi)

	//formatted print like C language
	// fmt.Printf("%f ", pi)

	//precision
	fmt.Printf("%.3f ", pi)

	randNum := 1

	fmt.Println(randNum)

	var numOne = 1.000
	var num99 = .999

	fmt.Println(numOne - num99)

	const pi2 float64 = 3.1456787

	var (
		varA = 2
		varB = 3
	)

	var myName string = "Faraz Murtaza"

	fmt.Println(myName, varA, varB)

	fmt.Println(myName + " is my name")

	var isOver40 bool = true

	fmt.Printf("%T is the data type, %t is for boolean", pi2, isOver40)

	// %d for integer and others just like C language

	// 3 logical operators
	fmt.Println("true $$ false = ", true && false)

	//for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//relational operators
	// == != < > <= >=

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	//if statements
	yourAge := 8

	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else {
		fmt.Println("You can not drive")
	}

	//switch statements
	switch yourAge {
	case 16:
		fmt.Println("You can drive")
	case 18:
		fmt.Println("You can vote")
	default:
		fmt.Println("nothing")
	}
}
