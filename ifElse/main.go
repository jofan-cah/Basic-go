package main

import "fmt"

// func main() {
// 	// Declare a variable of type int
// 	var grade  string
// 	score := 80

// 	age := 10
// 	agee := 11
// 	if age == 10 {
// 		fmt.Println("Variable is equal to 10")
// 	}

// 	if agee == 10 {
// 		fmt.Println("Variable is equal to 10")
		
// 	}else{
// 		fmt.Println("Variable is not equal to 10")
// 	}


// 	if score >= 90 && score <= 100 {
// 		grade = "A"
// 		fmt.Println("Excellent", grade)
// 	} else if score >= 80 && score < 90 {
// 		grade = "B"
// 		fmt.Println("Very Good",grade)
// 	} else if score >= 70 && score < 80 {
// 		grade = "C"
// 		fmt.Println("Good",grade)
// 	} else if score >= 60 && score < 70 {
// 		grade = "D"
// 		fmt.Println("Not Bad",grade)
// 	} else if score >= 50 && score < 60 {
// 		grade = "E"
// 		fmt.Println("Bad",grade)
// 	} else if score >= 0 && score < 50 {
// 		grade = "F"
// 		fmt.Println("Awful",grade)
// 	} else {

// 		fmt.Println("Invalid Score")
// 	}
// }


func main() {
	number := 2
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}
}