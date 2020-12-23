package main

import "fmt"

func printVariables(myArray [5]int, mySliceFull, mySlice []int) {
	fmt.Println("myArray: ", myArray)
	fmt.Println("mySliceFull: ", mySliceFull)
	fmt.Println("mySlice: ", mySlice)
	fmt.Println()
}

func changeVariablesValue(myArray [5]int) {
	myArray[0] = 9
}

func changeVariablesReference(myArray *[5]int) {
	myArray[0] = 9
}

func main() {
	myArray := [5]int{0, 1, 2, 3, 4}
	mySliceFull := myArray[:]
	mySlice := myArray[1:3]

	printVariables(myArray, mySliceFull, mySlice)

	myArray[2] = 9
	fmt.Println("Changed myArray[2] to 9")
	printVariables(myArray, mySliceFull, mySlice)

	mySlice[1] = 8
	fmt.Println("Changed mySlice[1] to 8")
	printVariables(myArray, mySliceFull, mySlice)

	mySlice = append(mySlice, 7, 7, 7)
	fmt.Println("Appended 777 to mySlice")
	printVariables(myArray, mySliceFull, mySlice)

	changeVariablesValue(myArray)
	fmt.Println("Changed value of myArray[0] in a function to 9 (pass by value)")
	printVariables(myArray, mySliceFull, mySlice)

	changeVariablesReference(&myArray)
	fmt.Println("Changed value of &myArray[0] in a function to 9 (pass by reference)")
	printVariables(myArray, mySliceFull, mySlice)

}
