package main

import "fmt"

func main() {

	// array are initated with var [size]Type{optional elements}
	// given array
	a := [5]int{1, 2, 3, 4, 5}

	// Rotate the array by
	d := 1

	// Print to know
	fmt.Println("Given Array:", a)
	fmt.Println("Array Length:", len(a))
	fmt.Println("Rotate Array By:", d)

	rotateArray(a, d)
}

func rotateArray(a [5]int, d int) {
	//go has only one looping construct https://tour.golang.org/flowcontrol/1
	// number of rotation,  initiated with value zero
	var loop int
	// for used as while loop
	for loop < d {

		// store the first element in temp var
		temp := a[0]
		var i int
		// Left rotate the remaning elements
		// store the first element in an temp var and start moving each element to the left by one position
		// len(array) will return the length of a given array
		//for used here as three components loops
		for i = 0; i < len(a)-1; i++ {
			a[i] = a[i+1]
		}
		// store the value fo temp var to the last position since at this point i will last position
		a[i] = temp

		//increment loop variable
		loop++
	}

	// print the result array
	fmt.Println("Rotated array:", a)
}
