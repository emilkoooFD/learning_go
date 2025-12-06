package main

import (
	"errors"
	"fmt"
)

func main() {

	transactions := []float64{}

	fmt.Println("___Учет транзакций___")

	for {
		transact, err := outputTransaction()

		// Если не правильно введены данные вывожу ошибку
		if err != nil {
			fmt.Println("сумма не может быть нулевой")
			continue
		}

		// Присвоение в slice введенных транзакций пользователя
		transactions = append(transactions, transact)
		fmt.Println("Ваши транзакции: ", transactions)

		// Проверка если false - то выходим из цикла
		if !checkRepeatTransaction() {
			break
		}
	}
	balance := calculateBalance(transactions)
	fmt.Println("Ваш итоговый баланс: ", balance)

}

// Сбор транзакций пользователя и вывод ошибки
func outputTransaction() (float64, error) {
	var transact float64
	fmt.Println("Введите сумму транзакции: ")
	fmt.Scan(&transact)
	if transact == 0 {
		return 0, errors.New("no_params_error")
	}
	return transact, nil

}

// Проверка на дальнейшие децствия и возврат bool значения
func checkRepeatTransaction() bool {
	var repeatTransact string

	fmt.Println("Хотите добавить еще транзакций? (д/н): ")
	fmt.Scan(&repeatTransact)

	return repeatTransact == "д"
}

// Функция расчета баланса
func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
