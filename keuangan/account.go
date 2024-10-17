package keuangan

import (
	"errors"
	"fmt"
)

type Account struct {
	Nama    string
	Email string
	Saldo   Saldo
}
var Arr []Account
func init() {
	fmt.Println("Mulai Buat Akun")
	Arr = append(Arr, Account{"Lumoshive", "lumos@mail.com", Saldo{0.0}})
	Arr = append(Arr, Account{"Academy", "academy@mail.com", Saldo{0.0}})
	fmt.Println("Selesai Buat Akun")
}
func CekAccounts(){
	fmt.Printf("%+v\n", Arr)
}
func ValidasiNamaKosong(nama string) error{
	if nama == "" {
		return errors.New("Nama Tidak Boleh Kosong")
	}
	return nil
}
func CreateNewAccount(nama string, address string, saldo float32) Account{
	errorValidasi := ValidasiNamaKosong(nama)
	if errorValidasi != nil {
		fmt.Println(errorValidasi)
		return Account{}
	}
	result := Account{nama, address, CreateNewSaldo(saldo)}
	Arr = append(Arr, result)
	fmt.Print("Akun berhasil ditambahkan ")
	CekAccounts()
	return result
}