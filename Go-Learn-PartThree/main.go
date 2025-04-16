package main

import "fmt"

// func main() {
// 	m := map[string]string{
// 		"Google": "https://www.google.com",
// 	} // Это map. Тип данных в котором есть ключ и значение. То же самое, что словарь в питоне

// 	fmt.Println(m["Google"])
// 	m["Yandex"] = "https://www.yandex.ru" // так добавляется элемент в map
// 	m["Facebook"] = "https://www.facebook.com"
// 	delete(m, "Facebook") // так удаляется элемент из map
// 	// Если удалить элемент, котоорого нет, то ничего не случится, никакой ошибки не будет выведено
// 	// Также если попытаться вывести элемент, которого нет, то вернется пустая строка

// 	for key, value := range m { // так перебирается map
// 		fmt.Println(key, value)
// 	}
// }

// Приложение для менеджмента закладок (с меню)

type stringMap = map[string]string // Таким образом можно объявить тип данных, чтобы не писать каждый раз долгие названия
// Это также называется alias или псевдоним

func main() {
	bookMarks := stringMap{}

	fmt.Println("Добро пожаловать в менеджер закладок!")

Menu: // Это label, чтобы можно было использовать break для выхода из бесконечного цикла
	for {
		userInput := getUserInput()
		// Можно было бы здесь написать, например, Switch и если написать в кейсе break Switch, то мы бы вышли из всего приложения
		switch userInput {
		case 1:
			fmt.Println("Ваши закладки:")
			showMarks(bookMarks)
		case 2:
			addMark(bookMarks)
			fmt.Println("Закладка добавлена!")
		case 3:
			deleteMark(bookMarks)
			fmt.Println("Закладка удалена!")
		case 4:
			fmt.Println("Выходим из приложения...")
			break Menu
		default:
			fmt.Println("Неверный ввод! Попробуйте снова.")
			break Menu
		}
	}
}

func getUserInput() int {
	var answer int
	fmt.Println("Выберите действие:")
	fmt.Println("1. Посмотреть все закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход из приложения")
	fmt.Scan(&answer)
	return answer
}

func showMarks(bookMarks stringMap) {
	if len(bookMarks) == 0 {
		fmt.Println("Ой, вы еще не добавляли закладок")
	}
	for key, value := range bookMarks {
		fmt.Println(key, ":", value)
	}
}

func addMark(bookMarks stringMap) {
	var key string
	var value string
	fmt.Print("Введите название закладки:")
	fmt.Scan(&key)
	fmt.Print("Введите ссылку:")
	fmt.Scan(&value)
	bookMarks[key] = value
}

func deleteMark(bookMarks stringMap) {
	var key string
	fmt.Println("Введите название закладки для удаления: ")
	fmt.Scan(&key)
	delete(bookMarks, key)
}
