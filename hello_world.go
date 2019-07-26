package main

import "fmt"

/* this is a multiline
comment */

func main() {
	// fmt.Println("Hello world")
	//
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

	for _, value := range favNums3 {
		fmt.Println(value)
	}

	//slice

	numSlice := []int{5, 4, 3, 2, 1}

	numSlice2 := numSlice[3:5] // index 3 to 5 but 5 not included

	fmt.Println("numSlice[0]= ", numSlice2[0])

	numSlice3 := make([]int, 5, 10)
	for k := 0; k < 5; k++ {
		fmt.Println(numSlice3[k])
	}
	copy(numSlice3, numSlice)
	// fmt.Println(numSlice3[0])

	numSlice3 = append(numSlice3, 0, -1)

	// fmt.Println(numSlice3[6])

	presAge := make(map[string]int)

	presAge["ABCDEFGH"] = 42

	fmt.Println(presAge["ABCDEFGH"])
	fmt.Println(presAge)
	fmt.Println(len(presAge))

	listNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("sum: ", addThemUp(listNums))

	fmt.Println(SubtractThem(1, 2, 3, 4, 5))

	num3 := 3

	doubleNum := func() int {
		num3 *= 2
		return num3
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	fmt.Println(factorial(4))

	defer printTwo() // executes at the very end of function it is enclosed in, in this case the main function
	printOne()
	printOne()
	printOne()
	printOne()

	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	demoPanic()

	xVal := 0
	changeXVal(xVal) // will not change value of xVal because it is just passing value and not reference(pointers will help to change values)
	fmt.Println("x = ", xVal)

	xVal2 := 0
	changeXValPointer(&xVal2)
	fmt.Println("x = ", xVal2)
	fmt.Println("Memory adress of xVal2 is ", &xVal2)

	//declaring a pointer variable
	yPtr := new(int)
	changeYValuePointer(yPtr)
	fmt.Println("value of Y is ", *yPtr) //* for the VALUE at yPtr
}

func changeYValuePointer(yptr *int) {
	*yptr = 100
}

func changeXValPointer(x *int) {
	*x = 2
}

func changeXVal(x int) {
	x = 2
}

// in case error occurs, it goes to panic which just prints out the error
func demoPanic() {
	defer func() {
		// fmt.Println(recover())
		fmt.Println(recover())
	}()

	panic("PANIC")

}

// in case an error occurs, program does not crash, it just ignores that value that caused the error
func safeDiv(num1, num2 int) int {
	defer func() {
		// fmt.Println(recover())
		recover()
	}()

	solution := num1 / num2
	return solution
}

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}
	return sum
}

func SubtractThem(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
