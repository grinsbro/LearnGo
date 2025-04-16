package weather

import (
	"Go-learn-part-five/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Тесты в Go
// Тесты делятся на юнит тесты (Тестирует лишь небольшой фрагмент кода),
// интеграционные (Позволяют тестировать взаимодействие между частями программы, как они взаимодействуют)
// и end-to-end (Позволяют отследить взаимодействие пользователя от начала и до конца использования программы)
// Пишу тест в пакете geo

var ErrWrongFormat = errors.New("WRONG_FORMAT")

// Создаю функцию для получения погоды
func GetWeather(geo geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City) // Преобразую строку в объект структуры URL. Это нужно, чтобы к URL можно было добавить полученные от пользователя данные
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	params := url.Values{} // Создаю словарь для параметров URL
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	} else {
		params.Add("format", fmt.Sprint(format)) // Добавляю параметр "format" со значением format, которое получено от пользователя
	}
	baseUrl.RawQuery = params.Encode() // Преобразую параметры в строку и добавляю к URL
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}
