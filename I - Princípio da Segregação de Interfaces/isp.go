package main

import "fmt"

/*

I - Princípio da Segregação de Interfaces (Interface Segregation Principle)

Clientes não devem ser forçados a depender de métodos que não usam.
Devemos criar interfaces mais específicas do que uma genérica.
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

func CheckingAccount(t Transaction) {
	fmt.Println("CheckingAccount")
	t.TransferPix()
	t.TransferTed()
	t.TransferBroker()
}

func SavingAccount(t Transaction) {
	fmt.Println("SavingAccount")
	t.TransferPix()
	t.TransferTed()

	// Quebra do Princípio
	// A assinatura t.TransferBroker fica disponivel mais ela não deve ser usado.
}

/*
Solução:

Vamos aplicar o princípio de segregação de interface à nossa
função SavingAccount, tornando ela mais específica em termos de seus requisitos.
Ela só precisa de algo que seja específico.

*/

type Payment2 struct {
	value float64
}

func (p Payment2) TransferPix2() string {
	return "transferPix..."
}

func (p Payment2) TransferTed2() string {
	return "transferTed..."
}

func (p Payment2) TransferBroker2() string {
	return "transferBroker..."
}

type Transaction2 interface {
	TransferPix() string
	TransferTed() string
}

type TransactionInvestment interface {
	TransferBroker() string
}

type InvestmentTransaction interface {
	Transaction2
	TransferBroker() string
}

func CheckingAccount2(i InvestmentTransaction) {
	fmt.Println("CheckingAccount")
	i.TransferPix()
	i.TransferTed()
	i.TransferBroker()
}

func SavingAccount2(t Transaction2) {
	fmt.Println("SavingAccount")
	t.TransferPix()
	t.TransferTed()
}
