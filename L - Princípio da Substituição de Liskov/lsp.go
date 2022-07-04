package main

import "fmt"

/*

L - Liskov Substitution Principle (Princípio da substituição de Liskov)

O princípio da substituição de Liskov declara que as subclasses
devem ser substituíveis por suas classes de base.

*/

type Payment struct {
	value float64
}

func (p Payment) TransferPix() string {
	return "transferPix..."
}

func (p Payment) TransferTed() string {
	return "transferTed..."
}

func (p Payment) TransferBroker() string {
	return "transferBroker..."
}

type Transaction interface {
	TransferPix() string
	TransferTed() string
	TransferBroker() string
}

func CheckingAccoun(t Transaction) {
	fmt.Println("CheckingAccount")
	t.TransferPix()
	t.TransferTed()
	t.TransferBroker()
}

func SavingAccoun(t Transaction) {
	fmt.Println("SavingAccount")
	t.TransferPix()
	t.TransferTed()

	// Uma conta poupança não pode transferir para um broker

	// Quebra do Princípio
	// SavingAccount não podem ser usadas como base

}

func main() {
	payment := Payment{value: 1200.00}

	CheckingAccoun(payment)
	SavingAccoun(payment)
}
