package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

// Это структура Encrypter, которая содержит одно поле - ключ для шифрования
type Encrypter struct {
	Key string
}

// В конструкторе записываю, чтобы при создании экземпляра структуры переменная key заполнялась из окружения
func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}

// Метод, который шифрует перед сохранением массив байт, который получается при записи логина, пароля и url
func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key)) // aes - встроенная библиотека в Go, которая позволяет шифровать данные. Принимает массив байт и возвращает зашифрованный блок и ошибку
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encrypter) Decrypt(encryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce := encryptedStr[:nonceSize]
	cipherText := encryptedStr[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
