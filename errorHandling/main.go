package main

import "fmt"

func logging() {
	fmt.Println("function telah selesai dijalankan")
	errorMsg := recover()
	fmt.Println("Error msg:", errorMsg)
}

func main() {
	fmt.Println("Hello World")
	// ha(0)
	janganPanicYa()
}

func ha(value int) {
	// Function logging tetap di invoke
	// walaupun terjadi error
	// HARUS DI ATAS YA GAES!
	defer logging()

	fmt.Println("Mulai function Ha")
	result := 10 / value
	fmt.Println("Result:", result)
}

func janganPanicYa() {
	defer logging()

	fmt.Println("Panik!!!!1")
	panic("JANGAN PANIK BOS!")
}
