package main

import (
	"PasswordManager-Grinsbro/account"
	"PasswordManager-Grinsbro/encrypter"
	"PasswordManager-Grinsbro/files"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите действие",
}

func main() {
	color.Green("Добро пожаловать в программу менеджера паролей")

	err := godotenv.Load()
	if err != nil {
		color.Red("Не удалось найти .env файл")
	}
	vault := account.NewVault(files.NewJsonDb("passwords.vault"), *encrypter.NewEncrypter())

Menu:
	for {
		answer := promptData(menuVariants...)
		menuFunc := menu[answer]
		if menuFunc == nil {
			color.Red("Выходим из программы...")
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	userLogin := promptData("Введите логин: ")
	userPassword := promptData("Введите пароль: ")
	userUrl := promptData("Введите адрес сайта: ")
	myAccount, err := account.NewAccount(userLogin, userPassword, userUrl)
	if err != nil {
		color.Red("Неправильный формат логина или URL")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, url string) bool {
		return strings.Contains(acc.Url, url)
	})
	if len(accounts) == 0 {
		color.Red("Не найдено аккаунтов с таким URL")
	}
	for _, i := range accounts {
		i.OutputData()
	}
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, login string) bool {
		return strings.Contains(acc.Login, login)
	})
	if len(accounts) == 0 {
		color.Red("Не найдено аккаунтов с таким логином")
	}
	for _, account := range accounts {
		account.OutputData()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByUrl(url)
	if isDeleted {
		color.Green("Аккаунт удален!")
	} else {
		color.Red("Не найдено аккаунтов с таким URL")
	}
}

func promptData(content ...string) string {
	for index, value := range content {
		if index == len(content)-1 {
			fmt.Printf("%v : ", value)
		} else {
			fmt.Println(value)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
