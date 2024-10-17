package main

import (
	"tugas/keuangan"
)
func main() {
	keuangan.CreateNewAccount("Fernando", "fernando@mail.com", 0)
	keuangan.Arr[2].Saldo.TambahSaldo(5000)
	keuangan.Arr[2].Saldo.KurangSaldo(1000)
	keuangan.CreateNewAccount("", "fernando@mail.com", 0)
	keuangan.Arr[2].Saldo.TambahSaldo(0)
	keuangan.Arr[2].Saldo.KurangSaldo(0)
}