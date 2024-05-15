package lsp

import "fmt"

/*

L - Liskov Substitution Principle (Princípio da substituição de Liskov)

O princípio da substituição de Liskov declara que as subclasses
devem ser substituíveis por suas classes de base.

*/

// Payment representa uma transação de pagamento genérica.
type Payment struct {
	value float64
}

// Métodos de Payment que implementam a interface Transaction.
func (p Payment) TransferPix() string {
	return "transferPix..."
}

func (p Payment) TransferTed() string {
	return "transferTed..."
}

func (p Payment) TransferBroker() string {
	return "transferBroker..."
}

// Transaction define operações que todas as classes de pagamento devem implementar.
type Transaction interface {
	TransferPix() string
	TransferTed() string
	TransferBroker() string
}

// Funções que assumem que todos os tipos de transações suportam todos os métodos.
func checkingAccount(t Transaction) {
	fmt.Println("CheckingAccount")
	t.TransferPix()
	t.TransferTed()
	t.TransferBroker()
}

func savingAccount(t Transaction) {
	fmt.Println("SavingAccount")
	t.TransferPix()
	t.TransferTed()
	// SavingAccount não deveria chamar TransferBroker.
}

/*
func main() {
	payment := Payment{value: 1200.00}
	CheckingAccount(payment)
	SavingAccount(payment)
}
*/

/*
	Para respeitar o LSP, é crucial reorganizar a estrutura e a interface
	de modo que subclasses ou implementações específicas não quebrem as expectativas
	da interface base
*/

// Transaction define operações básicas disponíveis para todos os pagamentos.
type NTransaction interface {
	TransferPix() string
	TransferTed() string
}

// BrokerTransaction estende NTransaction com operações adicionais para brokers.
type BrokerTransaction interface {
	NTransaction
	TransferBroker() string
}

// Payment implementa todas as transações, incluindo aquelas para brokers.
type NPayment struct {
	value float64
}

func (p NPayment) TransferPix() string {
	return "transferPix..."
}

func (p NPayment) TransferTed() string {
	return "transferTed..."
}

func (p NPayment) TransferBroker() string {
	return "transferBroker..."
}

// PaymentWithoutBroker implementa apenas a interface básica Transaction.
type PaymentWithoutBroker struct {
	Payment
}

// Remove a implementação de TransferBroker, deixando apenas as transações básicas.
func (p PaymentWithoutBroker) TransferBroker() string {
	return "operation not supported"
}

/*
func main() {
	payment := Payment{value: 1200.00}
	checkingAccount(payment) // Suporta todas as transações.

	paymentWithoutBroker := PaymentWithoutBroker{Payment{value: 1200.00}}
	savingAccount(paymentWithoutBroker) // Suporta apenas TransferPix e TransferTed.
}
*/
