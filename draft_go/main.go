package main

import "fmt"

// "math"

func main() {
	// math.Pow(x, y) - метод возводящий в степень (x - число возводимое в степень, y - степень в которую возводим)
	const USD_RUB = 90.0
	const USD_EUR = 1.1

	EURInRUB := USD_RUB * USD_EUR
	fmt.Println(EURInRUB)

}
