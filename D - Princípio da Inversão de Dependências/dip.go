package main

import (
	"fmt"
)

/*

D - Princípio da Inversão de Dependências (Dependency Inversion Principle)

Dependa de abstrações e não de implementações.

*/

type Payment struct {
	value float64
}

func (p Payment) TransferBroker() string {
	return "transferBroker..."
}

type InvestmentTransaction interface {
	TransferBroker() string
}

func CheckingAccount(i InvestmentTransaction) {
	fmt.Println("CheckingAccount")
	i.TransferBroker()
}

/*
Solução

Separamos nosso exemplo em camadas

O módulo de mais baixo nivel (Domain Layer)
assim como de mais alto nivel (Presenter layer, Infra Layer),
Todas as camadas devem depender de abstração.

obs: Inversão de Dependência não é igual a Injeção de Dependência.
perceba que apenas na Main realizamos o acomplamento e consequentemente a Injeção de Dependência

Nesse exemplo conseguimos ver todos os principios do SOLID e um pouco do Clean Architecture.


*/

//Domain Layer
type PaymentEntity struct {
	value float64
}

func (p PaymentEntity) getValue() float64 {
	return p.value
}

type ITransactionBrokerAccountUseCase interface {
	execute(p PaymentEntity)
}

//Data Layer
type TransactionService struct {
	apiXptoClient IApiXptoClient
}

func (service TransactionService) execute(p PaymentEntity) {
	payload := TransactionBrokerRequest{p.value}
	service.apiXptoClient.executeTransactionValue(payload)
}

//Infra Layer
type ClientHttp struct{}

func (c *ClientHttp) Post(params any) {
	fmt.Println("request...", params)
}

type IApiXptoClient interface {
	executeTransactionValue(r TransactionBrokerRequest)
}

type TransactionBrokerRequest struct {
	value float64
}

func (client ClientHttp) executeTransactionValue(r TransactionBrokerRequest) {
	client.Post(r)
}

//Presenter layer
type PresenterTransactionBroker struct {
	transactionBrokerAccount ITransactionBrokerAccountUseCase
}

func (p PresenterTransactionBroker) TransactionBrokerController(params PaymentEntity) {
	p.transactionBrokerAccount.execute(params)
}

func main() {
	clientHttp := ClientHttp{}
	service := TransactionService{&clientHttp}
	presenter := PresenterTransactionBroker{&service}
	transactionValue := PaymentEntity{1200}
	presenter.TransactionBrokerController(transactionValue)
}
