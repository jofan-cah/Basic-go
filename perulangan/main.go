package main

import "fmt"

func main() {
	// Create a new slice of integers

count := 10
for i := 0; i < count; i++ {
		fmt.Println(i)
}

	title := "Belajar Golang"

	for i := 0; i < len(title); i++ {
		fmt.Printf("character %c at index %d\n", title[i], i)
	}
}