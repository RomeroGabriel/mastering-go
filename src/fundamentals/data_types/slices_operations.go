package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Initial mySlice: ", mySlice)
	// Add an element to the end of the slice.
	fmt.Println("Appending an element --------->")
	mySliceAdded := append(mySlice, 70)
	fmt.Println("mySliceAdded: ", mySliceAdded)
	fmt.Println("mySlice without changes: ", mySlice)
	fmt.Println()
	// Remove an element from the slice by index.
	fmt.Println("Removing an element --------->")
	index := 3
	mySlice = append(mySlice[:index], mySlice[index+1:]...) // Remove element at index
	fmt.Println("mySlice after remove index 3: ", mySlice)
	fmt.Println("mySliceAdded: ", mySliceAdded)
	fmt.Println()

	// Insert an element at a specific position in the slice.
	fmt.Println("Insert at index --------->")
	index = 5
	mySliceAdded = append(
		mySliceAdded[:index],
		append([]int{1000}, mySliceAdded[index:]...)...) // Insert value at index
	fmt.Println("mySlice after remove: ", mySliceAdded)
	fmt.Println("mySliceAdded: ", mySliceAdded)
}
