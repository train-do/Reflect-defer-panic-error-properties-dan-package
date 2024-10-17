package keuangan

import (
	"errors"
	"fmt"
)

type Saldo struct {
	Balance float32
}
func ValidasiTransaksiNol(saldo float32) error{
	if saldo == 0 {
		return errors.New("Nilai tidak boleh 0")
	}
	return nil
}
func CreateNewSaldo(saldo float32) Saldo {
	return Saldo{saldo}
}

func (s *Saldo) TambahSaldo(saldo float32) {
	errorValidasi := ValidasiTransaksiNol(saldo)
	if errorValidasi != nil {
		fmt.Println(errorValidasi)
		return
	}
	s.Balance += saldo
	fmt.Print("Saldo berhasil ditambahkan ")
	CekAccounts()
}
func (s *Saldo) KurangSaldo(saldo float32) {
	errorValidasi := ValidasiTransaksiNol(saldo)
	if errorValidasi != nil {
		fmt.Println(errorValidasi)
		return
	}
	s.Balance -= saldo
	fmt.Print("Saldo berhasil dikurangi ")
	CekAccounts()
}