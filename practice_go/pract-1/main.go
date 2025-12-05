package main

import (
	"fmt"
	"math"
)

// Калькулятор индекса массы тела
func main() {

	for {
		fmt.Println("__ Калькулятор ИТМ __")
		// Присваиваю полученные значения от пользователя в переменные
		userHeight, userWeight := getUserInput()
		// Присваиваю переменной вычисленное значение ИМТ
		IMT := calculateIMT(userWeight, userHeight)
		// Вывод полученного результата
		outputResult(IMT)
		// Присваиваю переменной возвращенное булевое значение из функции проверки
		isRepeatCalculation := checkRepeatCalculation()
		// Проверяю, если имеет false - то прерываю цикл
		if !isRepeatCalculation {
			break
		}
	}

}

// Получение данных с консоли для расчета ИМТ
func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Println("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Println("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

// Расчет и возврат значения ИМТ
func calculateIMT(weight float64, height float64) float64 {
	IMT := weight / math.Pow(height/100, 2)
	return IMT
}

// Результат ИМТ и его вывод
func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела = %.0f", imt)
	fmt.Println(result)
	switch {
	case imt <= 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt <= 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt <= 25:
		fmt.Println("У вас нормальная масса тела")
	case imt <= 30:
		fmt.Println("У вас избыточная масса тела")
	case imt <= 35:
		fmt.Println("У вас 1-я степень ожирения")
	case imt <= 40:
		fmt.Println("У вас 2-я степень ожирения")
	default:
		fmt.Println("У вас 3-я степень ожирения")
	}
}

// Проверка нужно ли выполнить повторение вычисления
func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Printf("Хотите сделать расчет еще раз? (да/нет): ")
	fmt.Scan(&userChoice)
	return userChoice == "да"
}
