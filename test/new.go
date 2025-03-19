package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Create an initial map with some values
	originalMap := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}

	// Function to calculate the size of the map in bytes
	calculateMapSize := func(m map[string]int) uint64 {
		var totalSize uint64
		// Calculate the size of the map structure itself
		totalSize += uint64(unsafe.Sizeof(m))

		// Iterate over the map and calculate the size of each key and value
		for key, value := range m {
			totalSize += uint64(unsafe.Sizeof(key))   // Size of string key
			totalSize += uint64(unsafe.Sizeof(value)) // Size of int value
		}
		return totalSize
	}

	// Show the size of the map before modifying
	initialSize := calculateMapSize(originalMap)
	fmt.Printf("Size of map before modifications: %.6f MB\n", float64(initialSize)/1024/1024)

	// Increase the map size by adding a new key-value pair
	originalMap["date"] = 40
	afterAddSize := calculateMapSize(originalMap)
	fmt.Printf("\nSize of map after adding new entry: %.6f MB\n", float64(afterAddSize)/1024/1024)

	// Decrease the map size by removing a key-value pair
	delete(originalMap, "banana")
	afterRemoveSize := calculateMapSize(originalMap)
	fmt.Printf("\nSize of map after removing entry: %.6f MB\n", float64(afterRemoveSize)/1024/1024)

	// Add another new key-value pair
	originalMap["elderberry"] = 50
	afterSecondAddSize := calculateMapSize(originalMap)
	fmt.Printf("\nSize of map after adding second entry: %.6f MB\n", float64(afterSecondAddSize)/1024/1024)
}
