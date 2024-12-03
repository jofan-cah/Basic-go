package main

import "fmt"

// Struct Rectangle
type Rectangle struct {
	Width, Height float64
}

// Method untuk menghitung luas
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method untuk menghitung keliling
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Method untuk memperbarui dimensi
func (r *Rectangle) Resize(newWidth, newHeight float64) {
	r.Width = newWidth
	r.Height = newHeight
}

func main() {
	// Membuat instance Rectangle
	rect := Rectangle{Width: 10, Height: 5}

	// Menggunakan method Area dan Perimeter
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// Memperbarui dimensi menggunakan method Resize
	rect.Resize(20, 10)
	fmt.Printf("New Area: %.2f\n", rect.Area())
	fmt.Printf("New Perimeter: %.2f\n", rect.Perimeter())
}
