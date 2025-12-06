package main

import (
	"errors"
	"fmt"
)

// Калькулятор конвертации температуры
func main() {
	for {
		// Вывод переменных полученных с функции getData()
		fromScale, toScale, temperature, err := getData()
		if err != nil {
			fmt.Println("Не правильно введены данные")
			continue
		}
		// Присваиваю результат выполнения функции в переменную
		result := converterTemperature(fromScale, toScale, temperature)
		// Вывод в консоль полученных результатов
		fmt.Println(result)
		// Проверка ответа, если false - то завершаю цикл
		if !checkRepeat() {
			break
		}
	}
}

// Получение и возврат данных с консоли
func getData() (string, string, float64, error) {
	var fromScale string
	var toScale string
	var temperature float64
	fmt.Println("Из какой шкалы конвертировать?(c/f/k): ")
	fmt.Scan(&fromScale)
	fmt.Println("Введите температуру: ")
	fmt.Scan(&temperature)
	fmt.Println("В какую шкалу конвертировать?(c/f/k): ")
	fmt.Scan(&toScale)
	if (fromScale != "c" && fromScale != "f" && fromScale != "k") ||
		(toScale != "c" && toScale != "f" && toScale != "k") || temperature <= 0 {
		return "", "", 0, errors.New("no_params_error")
	}
	return fromScale, toScale, temperature, nil
}

// Проверка конвертации
func converterTemperature(of string, in string, temp float64) string {
	switch of {
	case "c":
		switch in {
		case "f":
			return fmt.Sprintf("%.2f°C = %.2f°F", temp, CinF(temp))
		case "k":
			return fmt.Sprintf("%.2f°C = %.2fK", temp, CinK(temp))
		case "c":
			return "Эта та же шкала"
		default:
			return "Неизвестная целевая шкала"
		}
	case "f":
		switch in {
		case "f":
			return "Эта та же шкала"
		case "k":
			return fmt.Sprintf("%.2f°F = %.2fK", temp, FinK(temp))
		case "c":
			return fmt.Sprintf("%.2f°F = %.2f°C", temp, FinC(temp))
		default:
			return "Неизвестная целевая шкала"
		}
	case "k":
		switch in {
		case "f":
			return fmt.Sprintf("%.2fK = %.2f°F", temp, KinF(temp))
		case "k":
			return "Эта та же шкала"
		case "c":
			return fmt.Sprintf("%.2fK = %.2f°C", temp, KinC(temp))
		default:
			return "Неизвестная целевая шкала"
		}
	default:
		return "Неизвестная исходная шкала"
	}

}

// Получение boolean значения для дальнейшей проверки
func checkRepeat() bool {
	repeat := ""
	fmt.Print("Хотите сделать рассчет еще раз? (д/н): ")
	fmt.Scan(&repeat)
	return repeat == "д"
}

// Функции с формулами конвертации

func CinF(celsia float64) float64 {
	return celsia*9/5 + 32
}

func FinC(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CinK(celsia float64) float64 {
	return celsia + 273.15
}

func KinC(kelvin float64) float64 {
	return kelvin - 273.15
}

func FinK(fahrenheit float64) float64 {
	celsius := FinC(fahrenheit)
	return CinK(celsius)

}

func KinF(kelvin float64) float64 {
	celsius := KinC(kelvin)
	return CinF(celsius)
}
