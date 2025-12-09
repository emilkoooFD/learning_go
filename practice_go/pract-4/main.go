package main

import "fmt"

type bookmarMap = map[string]string

func main() {
	// Начальный map с данными
	bookmarks := bookmarMap{}
	fmt.Println("___Приложение закладок___")
	for {
		// Вызов функции меню и присвоение в перменную вернувшего значения
		numberMenuSelection := getMenu()
		// Switch для понимания что нажали и какие действия нужно предпринять
		switch numberMenuSelection {
		case 1:
			fmt.Println("--Ваши закладки--")
			// Вызов функции показа закладок
			printBookmarks(bookmarks)
		case 2:
			fmt.Println("--Добавление закладки--")
			// Вызов функции добавления закладки
			addBookmark(bookmarks)
		case 3:
			fmt.Println("--Удаление закладки--")
			// Вызов функции удаления закладки
			deleteBookmark(bookmarks)
		case 4:
			return
		}
	}
}

// Функция показа меню и присвоение цифры в перменную
func getMenu() int {
	fmt.Println("Меню:")
	var numberMenu int
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&numberMenu)
	return numberMenu
}

// Функция вывода списка закладок в консоль
func printBookmarks(bm bookmarMap) {
	if len(bm) == 0 {
		fmt.Println("Ваш список закладок пуст! Добавьте их!")
	}
	for key, value := range bm {
		fmt.Println(key, ": ", value)
	}
}

// Функция добавления закладок
func addBookmark(bm bookmarMap) {
	var nameBookmark string
	var addressBookmarks string
	fmt.Print("Введите название (прекратить, наберите - n): ")
	fmt.Scan(&nameBookmark)
	fmt.Print("Введите адрес: ")
	fmt.Scan(&addressBookmarks)
	if nameBookmark == "n" || addressBookmarks == "n" {
		return
	}
	bm[nameBookmark] = addressBookmarks
}

// Функция удаления закладок
func deleteBookmark(bm bookmarMap) {
	var deleteNameBookmark string
	fmt.Scan(&deleteNameBookmark)
	delete(bm, deleteNameBookmark)
	fmt.Println("Вы удалили закладку!")
	fmt.Println(bm)
}
