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

	// var j int = 0 or j := "something"

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

	//arrays
	var favNums2 [5]float64
	favNums2[0] = 1.1
	favNums2[1] = 1.2
	favNums2[2] = 1.3
	favNums2[3] = 1.4
	favNums2[4] = 1.5

	fmt.Println(favNums2[3])

	favNums3 := [5]float64{1, 2, 3, 4, 5}

	for i, value := range favNums3 {
		fmt.Println(value, i)
	}

}
