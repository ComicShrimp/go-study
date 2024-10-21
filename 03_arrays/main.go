package main

import "fmt"

func main() {
	// Arrays
	arr := [...]int32{
		1,
		2,
		3,
	} // Array always with fixed length and fixed type and contigous in memory

	fmt.Println(arr)
	arr[1] = 20

	fmt.Println(arr[1])
	fmt.Println(arr[1:3])

	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	// Slices | Slices wraps an array, so under the hood is a "array with convinient functions calls"
	slic := []int32{10, 20, 30}

	fmt.Println(slic)
	fmt.Printf("The length of the slice: %v | The capacity %v\n", len(slic), cap(slic))

	slic = append(slic, 40)

	fmt.Println(slic)
	fmt.Printf("The length of the slice: %v | The capacity %v\n", len(slic), cap(slic))

	var slic2 []int32 = []int32{4, 5, 6}
	slic = append(
		slic,
		slic2...) // Append one slice with another slice | "..." is s spread operator
	fmt.Println(slic)

	var slic3 []int32 = make([]int32, 3, 8) // Create a slice specifying the length and the capacity of a slice
	fmt.Println(slic3)

	// Maps
	var mapping map[string]uint8 = make(map[string]uint8)
	fmt.Println(mapping)

	// Alternative initializing
	mapping2 := map[string]uint8{"test": 1, "test2": 2}
	fmt.Println(mapping2)
	fmt.Println(
		mapping2["Ops"],
	) // When acessing a key that doesnt exists, it return the default value for the type

	_, ok := mapping2["Ops"]

	if ok {
		fmt.Println("Value exists!")
	} else {
		fmt.Println("Value does not exists")
	}
}
