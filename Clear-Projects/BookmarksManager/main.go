package main

import (
	"fmt"

	"github.com/fatih/color"
)

type stringMap = map[string]string

func main() {
	color.Green("Доброе пожаловать в менеджер закладок")
	bookMarks := stringMap{}

Menu:
	for {
		response := getUserInput()
		switch response {
		case 1:
			fmt.Println("Ваши закладки: ")
			showBookMarks(bookMarks)
		case 2:
			addBookMark(bookMarks)
		case 3:
			deleteBookMark(bookMarks)
		default:
			color.Red("Выходим из приложения...")
			break Menu
		}
	}

}

func getUserInput() float64 {
	var answer float64
	fmt.Println("Выберите действие: ")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход из приложения")
	fmt.Scan(&answer)
	return answer
}

func showBookMarks(bookmarks stringMap) {
	if len(bookmarks) == 0 {
		color.Red("Ой, вы еще ничего не добавили...")
	}
	for _, value := range bookmarks {
		fmt.Println(value)
	}
}

func addBookMark(bookmarks stringMap) {
	var key string
	var value string
	color.Cyan("Введите название закладки: ")
	fmt.Scan(&key)
	color.Cyan("Введите ссылку: ")
	fmt.Scan(&value)
	bookmarks[key] = value
}

func deleteBookMark(bookmarks stringMap) {
	var answer string
	color.Red("Введите название закладки для удаления: ")
	fmt.Scan(&answer)
	delete(bookmarks, answer)
}
