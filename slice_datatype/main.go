package main

import "fmt"

func main() {

	// Direct declaration
	var values = [...]int{10, 20, 30, 40, 50, 60}

	// Indirect declaration
	var slices []int = values[:]

	var slice1 = values[4:6]
	var slice2 = values[:3]
	var slice3 = values[3:]
	var slice4 = values[:]
	var slice5 = values[:2]

	fmt.Println(slices)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	/*
		Function of slice
	*/
	fmt.Printf("slice1 length:%d\n", len(slice1))
	fmt.Printf("slice1 capacity:%d\n", cap(slice1))

	// This slice's capacity is still available

	fmt.Println("Appending 35 to slice5")
	fmt.Printf("slice5 cap:%d\n", cap(slice5))
	var newSlice5 = append(slice5, 35)

	fmt.Println(newSlice5)
	fmt.Println(values) // Won't change because capacity still available


	// This slice is out of capacity, so will built a new one
	fmt.Println("Appending 70 to slice1")
	fmt.Printf("slice1 cap:%d\n", cap(slice1))
	var newSlice1 = append(slice1, 70)

	fmt.Println(newSlice1)
	fmt.Println(values) // Will change because capacity not available

	// Slice can affect the array assigned
	var newValues = [...]string{"Lee", "'Apollo'", "Adama"}
	var sliceAffect = newValues[1:]
	sliceAffect[0] = "'Husker'"
	sliceAffect[1] = "Dima"

	fmt.Println(newValues)

	// 'Make' slice. The array
	// would be automatically created
	// make(slice_type, slice_length, slice_capacity)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Ensign"
	newSlice[1] = "William"
	// newSlice[2] = "Adama" <- Error, must use append()
	var newSlice2 = append(newSlice, "Adama")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// newSlice2 would still use the same array
	// because there is still capacity left
	newSlice2[0] = "Liutenant"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// 'Copy' slice.
	var starbuck = [...]string { "Kara", "'Starbuck'", "Thrace" }

	var fromSlice = starbuck[:]
	var toSlice = make([]string, len(fromSlice), cap(fromSlice))
	// Before copying
	fmt.Println(toSlice)

	copy(toSlice, fromSlice)

	// After copying
	fmt.Println(toSlice)

	// Stark difference between array and slice

	var thisArray = [...]int { 1, 2, 3, 4, 5}
	var thisSlice = []int { 1, 2, 3, 4, 5 }

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
