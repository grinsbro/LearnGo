package filemanagement

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
} // Это структура JsonDb, которая хранит в себе имя файла в виде строки. Она нужна для того, чтобы не передавать имя файла в каждую функцию, а просто использовать его в методах структуры JsonDb

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) { // Я переписал метод Read так, чтобы он был методом структуры JsonDb. Теперь он будет использовать не фиксированное имя, а имя файла, которое было передано в структуру JsonDb
	color.Green("Reading file")
	// file err := os.Open("test.txt") // Таким образом открывается файл, но таким образом файл читается по байтам
	data, err := os.ReadFile(db.filename) // Таким образом файл читается целиком. Это можно использовать, когда файл небольших размеров
	// Также я переписал функцию ReadFile так, чтобы она принимала имя файла из структуры JsonDb
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(data) // Это байтовый массив, который нужно преобразовать в строку
	// fmt.Println(string(data))
	return data, nil
}

func (db *JsonDb) Write(content []byte) { // Я также переписал метод Write так, чтобы он работал с с filename, который хранится в структуре JsonDb. Теперь он будет использовать не фиксированное имя, а имя файла, которое было передано в структуру JsonDb
	color.Blue("Writing file")
	file, err := os.Create(db.filename) // Таким образом создается файл. Используется встроенный пакет os для работы с файловой системой компьютера, а затем вызывается метод Create, который создает файл с указанным именем
	if err != nil {
		fmt.Println(err)
	}
	// _, err = file.WriteString(content) // Таким образом в файл записываются данные. В данном случае записывается строка, но также есть метод Writebyte, который записывает байты
	_, err = file.Write(content)
	defer file.Close() // Ключевое слово defer откладывает выполнение кода на конец стек фрейма. То есть то, что написано в defer, то выполнится самым последним
	// Также если использовать несколько defer, то последним выполнится самый первый defer
	if err != nil {
		fmt.Println(err)
		return
	}
	color.Green("Запись успешна")
}
