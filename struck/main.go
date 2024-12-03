package main

import "fmt"

type Studend struct {
	name     string
	age      int
	gender   string
	isActive bool
}

// func main() {
// 	// var s studend
// 	// s.name = "Rizky"
// 	// s.age = 20
// 	// s.gender = "Laki-Laki"
// 	// s.isActive = true
// 	// fmt.Println(s)

// 	user := Studend{
// 		name:     "Rizky",
// 		age:      20,
// 		gender:   "Laki-Laki",
// 		isActive: true,
// 	}
// 	fmt.Println(user)

// }

func main() {
	// Membuat instance struct
	person := Studend{
		name:     "Rizky",
		age:      20,
		gender:   "Laki-Laki",
		isActive: true,
	}

	// Memanggil fungsi dengan struct sebagai parameter
	greet(person)
}

// Fungsi menerima struct sebagai parameter
func greet(p Studend) {
	fmt.Printf("Hello, %s! You are %d years old, Kamu berjenis %s dan %t.\n", p.name, p.age, p.gender, p.isActive)
}
