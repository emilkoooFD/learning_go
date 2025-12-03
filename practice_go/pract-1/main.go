package main

import (
	"fmt"
	"math"
)

// Калькулятор индекс массы тела
func main() {
	var userHeight float64
	var userWeight float64
	fmt.Println("__ Калькулятор ИТМ __")
	fmt.Println("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight/100, 2)
	fmt.Printf("Ваш индекс массы тела = %.2f", IMT)
}
