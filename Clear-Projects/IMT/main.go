package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/fatih/color"
)

func main() {
	color.Green("Добро пожаловать в калькулятор индекса массы тела!")

	for {
		userHeight := promptData("Введите ваш рост: ")
		userWeight := promptData("Введите ваш вес: ")

		IMT, err := calculateIMT(userHeight, userWeight)
		if err != nil {
			panic("Не введены данные для расчета")
		}

		outputResult(IMT)

		var anotherOne string
		fmt.Println("Хотите сделать еще один расчет? [y/n]")
		fmt.Scan(&anotherOne)
		if anotherOne == "y" {
			continue
		} else {
			color.Red("Выходим из программы...")
			break
		}
	}
}

func promptData(prompt string) float64 {
	var answer float64
	fmt.Println(prompt)
	fmt.Scan(&answer)
	return answer
}

func calculateIMT(userHeight, userWeight float64) (float64, error) {
	const Pow = 2
	if userHeight <= 0 || userWeight <= 0 {
		color.Red("Некорректный ввод данных")
		return 0, errors.New("NO_PARAMS_ERROR")
	}
	IMT := userWeight / math.Pow(userHeight/100, Pow)
	return IMT, nil
}

func outputResult(IMT float64) {
	color.Cyan("Ваш индекс массы тела: %.2f", IMT)
	if IMT < 16 {
		color.Red("У вас сильный дефицит массы тела")
	} else if IMT < 18.5 {
		color.Yellow("У вас недостаточная масса тела")
	} else if IMT < 25 {
		color.Green("У вас нормальная масса тела")
	} else if IMT < 30 {
		color.Yellow("У вас избыточная масса тела")
	} else {
		color.Red("У вас ожирение")
	}

}
