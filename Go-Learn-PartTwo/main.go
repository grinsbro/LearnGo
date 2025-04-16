package main

import "fmt"

// Массивы
// Массивы это набор однотипных значений
// func main() {
// 	transactions := [6]int{5, 10, -7, 4, 5, 6}
// 	// В Go массивы создаются таким образом. Необязательно указывать длину массива, можно написать "..."
// 	// banks := [2]string{"Тинькофф", "Альфа"}

// 	// var currencies [3]string
// 	// currencies = [3]string{"USD", "EUR", "RUB"}

// 	// fmt.Println(transactions[1])
// 	// fmt.Println(banks[0])
// 	// fmt.Println(currencies[2])

// 	// partial := transactions[1:4] // Это срез или слайс. Срез это динамический массив, который может изменять свою длину
// 	// Также можно указать и без первого значения, и без последнего, чтобы получить срез от начала до нужного индекса или от нужного индекса до конца массива соответственно
// 	// В целом работа с массивами в Go очень похожа на работу с массивами в других языках программирования, но есть некоторые отличия
// 	// Например, в Go массивы передаются по значению, а не по ссылке. Это значит, что если передаешь массив в функцию, то передаешь его копию, а не ссылку на него

// 	transactionsNew := transactions // Это копия массива, а не ссылка на него
// 	transactionsNew[0] = 100        // Изменяем первый элемент массива, но это не изменит оригинальный массив
// 	// fmt.Println(transactions)
// 	// fmt.Println(transactionsNew)
// 	transactionsPartial := transactions[1:] // Слайс это ссылка на оригинальный массив, поэтому если мы изменим слайс, то изменится и оригинальный массив
// 	transactionsPartial[0] = 100            // Изменяем первый элемент среза, это изменит оригинальный массив
// 	// fmt.Println(transactions)    // Изменился оригинальный массив, потому что срез это ссылка на оригинальный массив
// 	// fmt.Println(transactionsPartial)
// 	transactionsNewPartial := transactionsPartial[:1]
// 	transactionsNewPartial[0] = 200 // Изменяем первый элемент среза, это изменит оригинальный массив
// 	// fmt.Println(transactions)
// 	// fmt.Println(transactionsPartial)                                  // Даже слайс слайса изменяет оригинальный массив. В целом логично, потому что ссылаются они изначально на один и тот же массив
// 	// fmt.Println(len(transactionsNewPartial), cap(transactionsNewPartial)) // len показывает длину массива в целом как и везде, а cap (capacity) показывает то, сколько еще значений можно добавить в массив. То есть cap показывает, сколько чисел может быть справа
// 	// В Go, например можно сделать так:

// 	transactionsNewPartial = transactionsNewPartial[:4] // Такое возможно, потому что у этого массива есть capacity добавить еще числа справа и несмотря на то, что мы сначала сделали срез по 1 индексу, мы все равно можем сделать срез по 4 индексу, потому что у этого массива есть capacity добавить еще числа справа
// 	fmt.Println(transactionsNewPartial)
// }

// Динамические массивы (слайсы)
// func main() {
// transactions := []int{0, 20, 35} // Фактически это слайс, потому что мы не указываем длину, но под капотом это массив
// newTransactions := append(transactions, 100) // append добавляет элемент в конец массива, но не изменяет оригинальный массив, а создает новый массив и возвращает его
// append добавляет capacity и невозможно применить его к массиву, потому что массивы имеют фиксированную длину и не могут быть изменены
// Например, если сделать так:
// temp := transactions
// transactions = append(transactions, 200)
// Несмотря на то, что temp тоже должен был поменяться, но это не так, потому что append создает новый массив и резервирует место в памяти как для нового элемента
// Получается, что теперь transactions имеет новую версию и будет содержать новые элементы, а temp будет ссылаться на старую версию transactions и будет содержать старые элементы
// fmt.Println(transactions)
// fmt.Println(temp)
// }

// Упражнение по работе с массивами и слайсами, где пользователь может вводить свои траты и прибыли, а программа будет выводить в конце массив

// func main() {
// 	fmt.Println("Добро пожаловать в программу учета финансов!")
// 	var transactions = []float64{}
// 	for {
// 		operation := scanTransaction()
// 		if operation == 0 {
// 			fmt.Printf("Вот список ваших транзакций: %v\n", transactions)
// 			break
// 		}
// 		transactions = append(transactions, operation)
// 		fmt.Println("Транзакция добавлена!")
// 	}
// }

// func scanTransaction() float64 {
// 	var operation float64
// 	fmt.Println("Введите транзакцию для учета, если хотите получить результат, введите 0:")
// 	fmt.Scan(&operation)
// 	return operation
// }

// Продолжение работы с массивами (циклы по массивам и слайсам)

// func main() {
// 	tr1 := []int{1, 2, 3}
// 	tr2 := []int{4, 5, 6}
// 	// tr1 = append(tr1, 7, 8, 9) // append принимает несколько аргументов для добавления в массив
// 	// Чтобы добавить к слайсу слайс, нужно использовать unpack, который записывается как ...
// 	// Например:
// 	tr1 = append(tr1, tr2...)

// 	// Цикл по массиву:
// 	for i, v := range tr1 { // range возвращает индекс и значение элемента массива. Нужно всегда испольлзовать два аргумента, если не нужно использовать индекс, то можно использовать "_" вместо него и наоборот
// 		fmt.Println(i, v)
// 	}

// 	fmt.Println("Добро пожаловать в программу учета финансов!")
// 	var transactions = []float64{}
// 	for {
// 		operation := scanTransaction()
// 		if operation == 0 {
// 			fmt.Printf("Вот список ваших транзакций: %v\n", transactions)
// 			break
// 		}
// 		transactions = append(transactions, operation)
// 		fmt.Println("Транзакция добавлена!")
// 	}
// }

// func scanTransaction() float64 {
// 	var operation float64
// 	fmt.Println("Введите транзакцию для учета, если хотите получить результат, введите 0:")
// 	fmt.Scan(&operation)
// 	return operation
// }

// Дополнение к приложению учета финансов. Добавляю вывод итогового баланса

func main() {
	fmt.Println("Добро пожаловать в программу учета финансов!")
	var transactions = []float64{}
	for {
		operation := scanTransaction()
		if operation == 0 {
			break
		}
		transactions = append(transactions, operation)
	}
	balance := calculateBalance(transactions)
	if balance > 0 {
		fmt.Printf("Ура, Вы в плюсе! Ваш баланс составляет: %.2f\n", balance)
	} else if balance < 0 {
		fmt.Printf("Увы, Вы в минусе! Ваш баланс составляет: %.2f\n", balance)
	} else {
		fmt.Println("Вы на нуле!")
	}
}

func scanTransaction() float64 {
	var operation float64
	fmt.Println("Введите транзакцию для учета, если хотите получить результат, введите 0:")
	fmt.Scan(&operation)
	fmt.Println("Транзакция добавлена!")
	return operation
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
