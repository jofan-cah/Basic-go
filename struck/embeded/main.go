package main

import "fmt"

// Struct Person untuk individu
type Person struct {
	Name     string
	Age      int
	Gender   string
	IsActive bool
}

// Struct Group untuk grup orang
type Group struct {
	GroupName string
	Members   []Person
}

func main() {
	// Membuat beberapa instance Person
	member1 := Person{
		Name:     "Rizky",
		Age:      20,
		Gender:   "Laki-Laki",
		IsActive: true,
	}

	member2 := Person{
		Name:     "Ayu",
		Age:      22,
		Gender:   "Perempuan",
		IsActive: false,
	}

	member3 := Person{
		Name:     "Budi",
		Age:      25,
		Gender:   "Laki-Laki",
		IsActive: true,
	}

	// Membuat instance Group
	group := Group{
		GroupName: "Belajar Golang",
		Members:   []Person{member1, member2, member3},
	}

	// Memanggil fungsi untuk menampilkan informasi grup
	displayGroupInfo(group)
}

// Fungsi untuk menampilkan informasi grup
func displayGroupInfo(g Group) {
	fmt.Printf("Group Name: %s\n", g.GroupName)
	fmt.Println("Members:")
	for _, member := range g.Members {
		fmt.Printf("- %s, Age: %d, Gender: %s, Active: %t\n", member.Name, member.Age, member.Gender, member.IsActive)
	}
}
