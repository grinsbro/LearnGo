package main

import (
	"fmt"
	"go-learn-part-four/account" // Таким образом импортируется пакет, который был создан. Всегда необходимо указывать имя модуля, который был создан в go.mod, а потом путь к папке
	"go-learn-part-four/encrypter"
	"go-learn-part-four/filemanagement"
	"go-learn-part-four/output"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// func main() {
// 	a := 5
// 	// pointerA := &a // Это указатель на переменную a. Сама по себе переменная a - int, и поэтому если мы ее используем в функции, и потом, например, перезаписываем в функции значение, то это не мутирует изначальное значение
// 	// С указателем значение будет мутировать, потому что указатель указывает на адрес в памяти
// 	double(&a)
// 	fmt.Println(a)
// 	// fmt.Println(pointerA)  // Здесь выведется адрес переменной в памяти, а не значение
// 	// fmt.Println(*pointerA) // Здесь мы разыменовываем указатель, чтобы получить значение переменной, на которую он указывает
// }

// func double(num *int) {
// 	*num = *num * 2
// } // Если необходимо проводить манипуляции со значением записанным в указатель, то нужно обязательно разыменовать указатель

// Написание функции для вывода реверсного массива
// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	reverse(&a)
// 	fmt.Println(a)
// }

// func reverse(arr *[4]int) {
// 	for index, value := range *arr {
// 		arr[len(arr)-1-index] = value
// 	}
// }

// type account struct {
// 	login    string
// 	password string
// 	url      string
// } // Это стракт или структура. Структура это набор полей, которые могут быть разного типа. Структуры это что-то похожее на классы в других языках программирования, но в Go нет наследования (В привычном смысле) и полиморфизма, поэтому структуры это просто набор полей

// // Наследование в Go реализуется через композицию.
// type accountWithTimeStamp struct {
// 	createdAt time.Time
// 	updatedAt time.Time
// 	account   // Это встраивание структуры. То есть при инстанциировании структуры accountWithTimeStamp, в ней будет доступна структура account
// }

// // Методы страктов
// // Методы в основном записываются рядом со страктами
// func (acc account) outputData() { // Чтобы объявить метод стракта, нужно указать его имя между объявлением функции и ее именем
// 	fmt.Println(acc.login, acc.password, acc.url)
// }

// // Переписываю функцию генерации пароля в метод структуры
// func (acc *account) generatePassword(n int) {
// 	password := make([]rune, n)
// 	for i := range password {
// 		password[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	acc.password = string(password)
// }

// // Это конструктор структуры. В отличие от других языков он записывается вне структуры. Здесь прописывается логика по созданию инстанций структуры
// func newAccount(login, password, urlString string) (*account, error) { // Общепринято называть конструкторы по имени структуры добавляя new в начале
// 	if login == "" {
// 		return nil, errors.New("login cannot be empty")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("invalid URL")
// 	}
// 	newAcc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
// 	if password == "" {
// 		newAcc.generatePassword(8)
// 	}
// 	return newAcc, nil
// }

// // Конструктор наследующей структуры:
// func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) { // Общепринято называть конструкторы по имени структуры добавляя new в начале
// 	if login == "" {
// 		return nil, errors.New("login cannot be empty")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("invalid URL")
// 	}
// 	newAcc := &accountWithTimeStamp{
// 		createdAt: time.Now(),
// 		updatedAt: time.Now(),
// 		account: account{
// 			login:    login,
// 			password: password,
// 			url:      urlString,
// 		}, // Так зыписывается логика создания инстанции структуры, которая наследует другую структуру
// 	}
// 	if password == "" {
// 		newAcc.generatePassword(8) // Такая запись все равно валидна, потому что мы инстанциировали структуру account внутри структуры accountWithTimeStamp, и поэтому у нас есть доступ к методу generatePassword
// 	}
// 	return newAcc, nil
// }

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
} // Map может хранить функции. В случае с меню мы можем сделать ключи опциями, который выбирает пользователь, а в значениях будут функции, которые соответствуют опциям.
// Важно, что функции должны принимать одинаковые аргументы и возвращать одинаковые значения. В данном случае функции принимают указатель на структуру VaultWithDb, а возвращают ничего. То есть у них одинаковая сигнатура
// Также важно не вызывать функцию при записи в map, а просто записать ее

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите действие",
}

func main() {

	color.Green("Добро пожаловать в программу менеджера паролей!")

	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти .env файл")
	}
	vault := account.NewVault(filemanagement.NewJsonDb("data.vault"), *encrypter.NewEncrypter()) // Создаем новый экземпляр структуры VaultWithDb, которая хранит аккаунты. Внутри структуры Vault создается экземпляр структуры JsonDb, которая отвечает за работу с файлом. Внутри структуры JsonDb создается файл data.json, который будет хранить данные аккаунтов в формате JSON
	// Переписал код так, чтобы при создании экземляра VaultWithDb принимался еще энкриптер
Menu:
	for {
		answer := promptData(menuVariants...) // Можно вынести меню в отдельный массив строк, но чтобы разные элементы были раскиданы по разным аргументам функции, нужно использовать ... в аргументах функции. Это называется распаковка массива. То есть мы распаковываем массив строк в аргументы функции, чтобы они были разными аргументами

		// Создаю переменную, которая будет вызывать элемент мапы в зависимости от выбранного ответа
		menuFunc := menu[answer]
		if menuFunc == nil {
			color.Red("Выходим из программы...")
			break Menu
		}
		menuFunc(vault)

		// switch answer {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	color.Red("Выходим из программы...")
		// 	break Menu
		// }

	}

	// filemanagement.WriteFile("Привет! Я файл", "test.txt")
	// filemanagement.ReadFile()

	// str := []rune("Привет!)") // Это рунный массив. Руна это по alias для int32, который используется для хранения символов в Go. То есть, когда мы хотим пройтись for по строке, то изначально все символы будут переведены в unicode, и если мы хотим получить сами символы, то нужно вызывать функцию string()
	// for _, ch := range string(str) {
	// 	fmt.Println(ch, string(ch))
	// }

	// userLogin := promptData("Введите логин: ")
	// userPassword := promptData("Введите пароль: ")
	// fmt.Println("Введите сколько символов вы хотите в пароле: ")
	// var userPasswordLength int
	// fmt.Scan(&userPasswordLength)
	// userUrl := promptData("Введите URL: ")

	// account1 := account{
	// 	login,
	// 	password,
	// 	url,
	// } // Можно инстанциировать структуру таким образом, но тогда порядок очень важен. Нужно передавать все в том же порядке, что и в структуре, иначе возникнет путаница
	// myAccount := account{
	// 	login: userLogin,
	// 	url:   userUrl,
	// } // Если инстранциировать структуру таким образом, то можно записывать в любом порядке
	// При такой записи можно даже пропустить одно значение, и тогда оно будет пустым. Если пропустить какую-либо из переменных в первом случае, то выпадет ошибка

	// myAccount, err := newAccount(userLogin, userPassword, userUrl) // Так можно инстанциировать структуру, используя конструктор. Такой метод нужен, если есть сложная система валидации данных и тд.
	// if err != nil {
	// 	fmt.Println("Неверный формат URL или Логина")
	// 	return
	// }
	// // myAccount.generatePassword(userPasswordLength)
	// fmt.Println("Ваш новый пароль: ", myAccount.password)
	// myAccount.outputData() // Чтобы вызвать метод структуры, нужно сначала инстанциировать структуру, а потом вызвать метод структуры. Внутри метода можно использовать поля структуры, к которой он относится. То есть, если мы вызываем метод структуры account, то внутри метода мы можем использовать поля этой структуры

	// Инстанциирую новую структуру:
	// myAccount, err := account.NewAccount(userLogin, userPassword, userUrl)
	// Таким образом импортируется метод из пакета. Также как со встроенными пакетами
	// if err != nil {
	// 	fmt.Println("Неверный формат URL или Логина")
	// 	return
	// }

	// myAccount.OutputData()

	// outputPassword(&myAccount)

}

// func getMenu() int {
// 	color.Cyan("Выберите действие:")
// 	var answer int
// 	fmt.Println("1. Создать аккаунт")
// 	fmt.Println("2. Найти аккаунт")
// 	fmt.Println("3. Удалить аккаунт")
// 	fmt.Println("4. Выход")
// 	fmt.Scanln(&answer)
// 	return answer
// }

func createAccount(vault *account.VaultWithDb) {
	userLogin := promptData("Введите логин")
	userPassword := promptData("Введите пароль")
	userUrl := promptData("Введите URL")
	myAccount, err := account.NewAccount(userLogin, userPassword, userUrl)
	if err != nil {
		output.PrintError("Неверный формат URL или Логина") // Добавил вывод ошибки из отдельного пакета
		return
	}
	vault.AddAccount(*myAccount)
	// data, err := vault.ToBytes()
	// if err != nil {
	// 	color.Red("Ошибка обработки данных")
	// 	return
	// }
	// filemanagement.WriteFile(data, "data.json")
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, url string) bool {
		return strings.Contains(acc.Url, url)
	})
	if len(accounts) == 0 {
		color.Red("Не найдено аккаунтов с таким URL")
	}

	for _, account := range accounts {
		account.OutputData()
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
		color.Green("Аккаунт удален")
	} else {
		output.PrintError("Аккаунт не найден")
	}
}

// Перепишу функцию, чтобы она была дженериком и принимала слайс любого типа (В итоге убрал дженерик)
func promptData(prompt ...string) string { // Чтобы функция принимала любое количество аргументов, нужно указать ... в аргументах функции
	for index, value := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v : ", value)
		} else {
			fmt.Println(value)
		}
	}

	// fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}

// func outputPassword(acc *account) {
// 	fmt.Println(acc.login, acc.password, acc.url)
// }

// func generatePassword(n int) string {
// 	password := make([]rune, n)
// 	for i := range password {
// 		password[i] = letterRunes[rand.IntN(len(letterRunes))]
// 	}
// 	return string(password)
// }

// Функция может возвращать другую функцию
// В данном случае происходит замыкание. То есть функция menuCounter возвращает другую функцию, которая будет использовать переменную i, которая была объявлена в родительской функции menuCounter. То есть переменная i будет доступна только внутри функции menuCounter и функции, которую она возвращает
// func menuCounter() func() {
// 	i := 0
// 	return func() {
// 		i++
// 		fmt.Println(i)
// 	}
// }
