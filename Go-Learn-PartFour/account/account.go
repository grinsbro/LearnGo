package account

// Если необходимо вынести какой-то функционал в отдельный пакет, то необходимо в корне проекта создать папку, которая будет называться также как и пакет, и в ней создать файл

import (
	"errors"

	// "fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color" // Это импорт стороннего пакета. Перед этим необходимо его установить, прописав в терминале go get и путь до пакета. После этого в файле go.mod появятся зависимости, а также появится go.sum файл, в котором будут указаны версии зависимостей
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*")

type Account struct {
	Login     string    `json:"login"` // Это тэг, который позволяет указать, как будет называться поле в определенном формате. Тут указан json, но также можно указать и для xml, тогда запись будет `json:"login xml:"login` Теги отделяются пробелом
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type accountWithTimeStamp struct {
// 	createdAt time.Time
// 	updatedAt time.Time
// 	Account   // Это встраивание структуры. То есть при инстанциировании структуры accountWithTimeStamp, в ней будет доступна структура account
// }

func (acc *Account) OutputData() { // Чтобы объявить метод стракта, нужно указать его имя между объявлением функции и ее именем
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
	// fmt.Println(acc.login, acc.password, acc.url)
}

// func (acc *Account) ToBytes() ([]byte, error) {
// 	file, err := json.Marshal(acc) // Marshal это метод, который преобразует структуру в байтовый массив. В данном случае мы преобразуем структуру account в json формат. Если бы мы использовали MarshalIndent, то у нас был бы отформатированный json
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }

func (acc *Account) generatePassword(n int) {
	password := make([]rune, n)
	for i := range password {
		password[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(password)
}

func NewAccount(login, password, urlString string) (*Account, error) { // Общепринято называть конструкторы по имени структуры добавляя new в начале
	if login == "" {
		return nil, errors.New("login cannot be empty")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("invalid URL")
	}
	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(8)
	}
	return newAcc, nil
}

// Чтобы экспортировать функцию, ее название должно быть с большой буквы
// func NewAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) { // Общепринято называть конструкторы по имени структуры добавляя new в начале
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
// 		Account: Account{
// 			login:    login,
// 			password: password,
// 			url:      urlString,
// 		}, // Так зыписывается логика создания инстанции структуры, которая наследует другую структуру
// 	}
// 	// field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login") // reflect это встроенная библиотека Go, которая позволяет работать с рефлексией. В данном случае мы получаем поле login структуры accountWithTimeStamp. Метод Elem() возвращает указатель на структуру, а метод FieldByName() возвращает поле структуры по имени. Второй параметр - это ошибка
// 	// fmt.Println(string(field.Tag)) // Здесь выводится тэг для поля login, который мы записали в field
// 	if password == "" {
// 		newAcc.generatePassword(8) // Такая запись все равно валидна, потому что мы инстанциировали структуру account внутри структуры accountWithTimeStamp, и поэтому у нас есть доступ к методу generatePassword
// 	}
// 	return newAcc, nil
// }
