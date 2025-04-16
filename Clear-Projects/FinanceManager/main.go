package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	color.Green("Добро пожаловать в менеджер финансов!")
	var transactions = []float64{}
	for {
		operation := getData()
		if operation == 0 {
			break
		}
		transactions = append(transactions, operation)
	}
	balance := finalBalance(transactions)
	if balance > 0 {
		color.Green("Ура вы в плюсе!")
		color.Cyan("Ваш баланс: %.2f", balance)
	} else if balance < 0 {
		color.Red("О нет! Вы в минусе, нужно что-то делать")
		color.Cyan("Ваш баланс: %.2f", balance)
	} else {
		color.Yellow("Вы на нуле :(")
	}

}

func getData() float64 {
	var operation float64
	color.Yellow("Введите операцию для учета. Если хотите выйти введите 0 или n")
	fmt.Scan(&operation)
	return operation
}

func finalBalance(tr []float64) float64 {
	sum := 0.0
	for _, value := range tr {
		sum += value
	}
	return sum
}
