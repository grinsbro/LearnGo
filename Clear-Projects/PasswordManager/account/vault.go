package account

import (
	"PasswordManager-Grinsbro/encrypter"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account
	UpdatedAt time.Time
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red("Не удалось разобрать JSON файл")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, v := range vault.Accounts {
		isMatched := checker(v, str)
		if isMatched {
			accounts = append(accounts, v)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false

	for _, v := range vault.Accounts {
		isMatched := strings.Contains(v.Url, url)
		if !isMatched {
			accounts = append(accounts, v)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		color.Red("Ошибка преобразования в JSON формат")
	}
	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	encData := vault.enc.Encrypt(data)
	if err != nil {
		color.Red("Ошибка обработки данных")
	}
	vault.db.Write(encData)
}
