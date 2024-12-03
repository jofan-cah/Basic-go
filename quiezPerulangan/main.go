package main

import (
	"fmt"
)

func main() {
	title := "Belajar Golang"

	for i := 0; i < len(title); i++ {	
		fmt.Printf("character %c at index %d\n", title[i], i)
	}

	fmt.Println("-----------")
	fmt.Print("\n")
	fmt.Println("-----------")

	for i := 0; i < len(title); i++ {
		if title[i] % 2 == 0	{
			fmt.Printf("character %c at index %d\n", title[i], i)
		}
	}

	fmt.Println("-----------")
	fmt.Print("\n")
	fmt.Println("-----------")

	for index, value := range title {
		valueString := string(value)
		if valueString == "a" || valueString == "i" || valueString == "u" || valueString == "e" || valueString == "o" {
			fmt.Printf("character %c at index %d\n", value, index)
			
		}
	}

	fmt.Println("-----------")
	fmt.Print("\n")
	fmt.Println("-----------")

}