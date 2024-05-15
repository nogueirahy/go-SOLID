package ocp

import (
	"fmt"
)

/*
	O — Open-Closed Principle (Princípio Aberto-Fechado)

	Uma class/função/componente, devem estar abertos para extensão,
	mas fechados para modificação,

	Uma class/função/componente não podemos alterar, mais sim extender.

*/

// Payment representa a forma de pagamento mas viola o princípio OCP.
type Payment struct {
	value    float64
	cashback float64 // Inclusão inadequada segundo OCP, pois modifica a estrutura.
}

// Métodos que modificam diretamente a estrutura Payment
func (p *Payment) payCreditCard(value float64) {
	fmt.Println("payCreditCredit...")
}

func (p *Payment) payTicket() {
	fmt.Println("payTicket...")
}

// payCrypto adiciona funcionalidade de forma invasiva, violando OCP.
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

// NPayment serve como a base para todos os tipos de pagamento, fechado para modificação.
type NPayment struct {
	value float64
}

// Interfaces e métodos que executam funcionalidades.
type PaymentCreditCard struct {
	NPayment
}

func (p PaymentCreditCard) execute() {
	p.value = 100.0
	fmt.Println("PayCreditCard...", p.value)
}

type PaymentTicket struct {
	NPayment
}

func (p PaymentTicket) execute() {
	p.value = 23.0
	fmt.Println("PayTicket...", p.value)
}

type PaymentCrypto struct {
	NPayment
	cashback float64
}

func (p PaymentCrypto) execute() {
	p.value = 250.0
	p.cashback = 0.80
	fmt.Println("PayCrypto...", p.value)
	fmt.Println("PayCrypto... Cashback", p.cashback)
}

/*
func main() {
	var paymentCreditCard PaymentCreditCard
	paymentCreditCard.execute()

	var paymenyTicket PaymentTicket
	paymenyTicket.execute()

	var paymenyCrypto PaymentCrypto
	paymenyCrypto.execute()
}
*/
