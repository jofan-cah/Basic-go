package main

import "fmt"

func main() {
	// Slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // [1 2 3 4 5]
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 5
	// Slice with initial capacity
	slice2 := make([]int, 0, 10)
	fmt.Println(slice2) // []
	fmt.Println(len(slice2)) // 0
	fmt.Println(cap(slice2)) // 10
	// len itu apa  ?

	// cap itu apa ?


	// append
	slice2 = append(slice2, 1, 2, 3)
	fmt.Println(slice2) // [1 2 3]

	// copy
	slice3 := make([]int, 5)
	copy(slice3, slice)
	fmt.Println(slice3) // [1 2 3 4 5]


	// slice pake language programer
	var leanguague  [] string
	leanguague = append(leanguague, "Java", "Python")
	leanguague = append(leanguague,  "Go", "Rubby", "PHP")
	fmt.Println(leanguague) // [Java Python Go Rubby PHP]
	}
