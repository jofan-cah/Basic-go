package main

import "fmt"


func main()  {
	// Membuat Array
	// var nama [5]int
	// var nama [5]bool
	// var nama = [5]string{"Rizky", "Jofan", "Fathurahman", "Dea", "Milla"}
	// var leanguague = [5]string{"Java", "Python", "Go", "Rubby", "PHP"}
	var leanguague = [...]string{"Java", "Python", "Go", "Rubby", "PHP"}

for index := 0; index < len(leanguague); index++ {
	fmt.Println(leanguague[index])
}
	fmt.Println(leanguague)
}
