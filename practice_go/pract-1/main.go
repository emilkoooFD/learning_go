package main

import (
	"fmt"
	"math"
)

// Калькулятор ИМТ
func main() {
	var userHeight float64
	var userWeight float64
	fmt.Println("__ Калькулятор ИТМ __")
	fmt.Println("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight, 2)
	fmt.Println("Ваш ИМТ =", IMT)
}
