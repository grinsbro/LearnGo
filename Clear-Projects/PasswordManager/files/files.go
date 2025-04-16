package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	color.Green("Читаем файл с данными...")
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	color.Cyan("Записываем данные...")
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Print(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Print(err)
	}
	color.Green("Данные успешно записаны!")
}
