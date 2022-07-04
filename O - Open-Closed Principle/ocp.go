package main

import (
	"fmt"
)

/*
	O — Open-Closed Principle (Princípio Aberto-Fechado)

	Uma class/função/componente, devem estar abertos para extensão,
	mas fechados para modificação,

	Uma class/função/componente não podemos alterar, mais sim extender.

*/

type Payment struct {
	value    float64
	cashback float64 // Quebra do Princípio adicionando uma nova prop a struct.

}

func (p *Payment) payCreditCard(value float64) {
	fmt.Println("payCreditCard...")
}

func (p *Payment) payTicket() {
	fmt.Println("payTicket...")

}

// Nova demanda de negócio, precisamos adiconar pagamento
// para cripto e adicionar cashback

// Quebra do Princípio
// Estamos gerando um novo método para lugares onde não usa esse meio de pagamento...
func (p *Payment) payCrypto() {
	p.cashback = 5
	fmt.Println("payCrypto...")
}

/*
Solução:

	Devemos extender o funcionamento da nossa struct Payment, porém sem modificar
	a struct existente.

	Analisamos o contexto principal(Pagamento) abstraindo-os para uma interface.
	Devemos focar muito bem na abstração, para facilitar o conceito de aberto
	para extensão...

*/

type Payment2 struct {
	value float64
}

type PaymentCreditCard struct {
	Payment2
}

func (p PaymentCreditCard) execute() {
	p.value = 100.0
	fmt.Println("PayCreditCard...", p.value)
}

type PaymentTicket struct {
	Payment2
}

func (p PaymentTicket) execute() {
	p.value = 23.0
	fmt.Println("PayTicket...", p.value)
}

type PaymentCrypto struct {
	Payment2
	cashback float64
}

func (p PaymentCrypto) execute() {
	p.value = 250.0
	p.cashback = 0.80
	fmt.Println("PayCrypto...", p.value)
	fmt.Println("PayCrypto... Cashback", p.cashback)

}

func main() {
	var paymentCreditCard PaymentCreditCard
	paymentCreditCard.execute()

	var paymenyTicket PaymentTicket
	paymenyTicket.execute()

	var paymenyCrypto PaymentCrypto
	paymenyCrypto.execute()
}
