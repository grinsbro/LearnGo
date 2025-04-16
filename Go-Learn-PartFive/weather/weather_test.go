package weather_test

import (
	"Go-learn-part-five/geo"
	"Go-learn-part-five/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Arrange
	expectedCity := "London"
	geo := geo.GeoData{
		City: expectedCity,
	}
	format := 3

	// Act
	result, _ := weather.GetWeather(geo, format)

	// Assert
	if !strings.Contains(result, expectedCity) {
		t.Errorf("Ожидалось %v, получено %v", expectedCity, result)
	}
}

// Группы тестов нужны для того, чтобы проверять разные условия в тестах и не переписывать тест каждый раз под определенное значение
var testCases = []struct { // Для этого объявляется структура, которая будет иметь имя и то значение, которое необходимо проверить
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Negative format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	// Для того, чтобы применить группу тестов нужно пройтись for по экземплярам ранее созданной структуры:
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // Вызываю начало теста с помощью t.Run() и передаю в аргументы название теста, а также анонимную функцию, которая проводит тестирование
			// Arrange
			geo := geo.GeoData{
				City: "London",
			}
			// Теперь не нужно объявлять формат отдельно
			// Act
			_, err := weather.GetWeather(geo, tc.format) // формат берется из структуры для каждого тест кейса

			// Assert
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrWrongFormat, err)
			}
		})
	}
}
