package main

import "fmt"

// main
// func main() {
// 	sentence := printMyresult("Saya Sedang")
// 	println(sentence)
// }

// func printMyresult(sentence string) string  {
// 	newSentence := sentence + " Belajar Golang"
// 	return newSentence

// }

// func main() {
// 	result := add(1, 2)
// 	println(result)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// Function Mutiple return ?
func calculate(a int, b int) (int, int) {
    sum := a + b
    product := a * b
    return sum, product
}

// func main() {
//     // Calling the function and capturing the returned values
//     sum, product := calculate(3, 4)
    
//     // Printing the results
//     fmt.Println("Sum:", sum)
//     fmt.Println("Product:", product)
// }

//function dengan predefined value ?
// func main() {
//     // Calling the function and capturing the returned values
//     sum, _ := calculate(3, 4)
		
// 		// Printing the results
// 		fmt.Println("Sum:", sum)
// }


type GreetParams struct {
    Name    string
    Greeting string
}

func greet(params GreetParams) string {
    if params.Name == "" {
        params.Name = "Guest" // Default value for Name
    }
    if params.Greeting == "" {
        params.Greeting = "Hello" // Default value for Greeting
    }
    return fmt.Sprintf("%s, %s!", params.Greeting, params.Name)
}

func main() {
    // Calling the function without arguments
    fmt.Println(greet(GreetParams{})) // Output: Hello, Guest!

    // Calling the function with one argument
    fmt.Println(greet(GreetParams{Name: "Alice"})) // Output: Hello, Alice!

    // Calling the function with both arguments
    fmt.Println(greet(GreetParams{Name: "Bob", Greeting: "Hi"})) // Output: Hi, Bob!
}
