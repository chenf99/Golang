package main

import "fmt"

func main() {
	s := []int{50, 12, 13}
	printSlice(s)
	
	s = s[:0]
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)
	
	s = s[:3]
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
