package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Создаю структуру, в которой сожержится город
type GeoData struct {
	City string `json:"city"` // Добавляю тэг, чтобы при парсинге из данных по ip подставлялось нужное поле
}

// Добавляю структуру, чтобы проверять существует ли такой город или нет
type CityPopulationResponse struct {
	Error bool `json:"error"` // Тэг также для подставления нужного поля из json файла
}

// Перепишу код, чтобы использовать переменные ошибок в тестах
var ErrorNoCity = errors.New("NO_CITY") // Переменные, где вызываются ошибки должны начинаться с Err или Error, потому что иначе компилятор будет подсвечивать эту переменную
var ErrorNot200 = errors.New("NOT200")

// Функция, которая проверяет значение полученное от пользователя. Если значение передано не было, то выполняется GET запрос на сайт, который по ip вычисляет город
func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city) // Добавляю также проверку, что если полльзователь ввел город, то такой город должен существовать
		if !isCity {
			return nil, ErrorNoCity
		}
		return &GeoData{ // Если все ок, то создаю экземпляр структуры с переданным городом
			City: city,
		}, nil
	}
	response, err := http.Get("https://ipapi.co/json/") // Делаю get запрос, чтобы получить JSON данные по айпи
	// http встроенная в Go структура, которая позволяет отправлять запросы
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		fmt.Println(response.StatusCode)
		return nil, ErrorNot200 // Если статус код не равен 200, то возвращаю ошибку
	}
	defer response.Body.Close()            // Добавляю закрытие Body в конце функции, чтобы не возникло утечки памяти
	body, err := io.ReadAll(response.Body) // если все проверки были пройдены, то читаю переменную Body у response, потому что response это экземпляр структуры http
	if err != nil {
		return nil, err
	}
	var geo GeoData            // Создаю переменную geo - экземпляр структуры GeoData
	json.Unmarshal(body, &geo) // Разбираю JSON формат записанный в body и записываю в geo
	return &geo, nil
}

// Создаю функцию, которая проверяет существует ли введенный пользователем город
func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{ // Преобразую город, который передал пользователь в JSON формат. Перед этим прямо в аргументах Marshal создаю мапу с этим городом
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody)) // Метод POST принимает json формат
	// Метод POST принимает 3 аргумента - адрес, куда нужно сделать запрос, формат данных, которые будут передаваться, а также данные, которые будут переданы в формате среза байт
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // Читаю ответ после запроса и записываю в переменную body
	if err != nil {
		return false
	}
	var populationResponse CityPopulationResponse // Создаю переменную - экземпляр структуры для проверки существования города для "распаковки" JSON файла
	json.Unmarshal(body, &populationResponse)     // "распаковываю" ответ в новую переменную
	return !populationResponse.Error              // Возвращаю переменную
}
