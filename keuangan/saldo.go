package keuangan

import "fmt"

type Saldo struct {
	Balance float32
}

func CreateNewSaldo(saldo float32) Saldo {
	return Saldo{saldo}
}

func (s *Saldo) TambahSaldo(saldo float32) {
	s.Balance += saldo
	fmt.Print("Saldo berhasil ditambahkan ")
	CekAccounts()
}
func (s *Saldo) KurangSaldo(saldo float32) {
	s.Balance -= saldo
	fmt.Print("Saldo berhasil dikurangi ")
	CekAccounts()
}