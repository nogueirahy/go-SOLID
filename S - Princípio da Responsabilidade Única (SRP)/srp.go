package srp

/*

Single Responsiblity Principle (Princípio da responsabilidade única)

Uma class/função/package/componente, deve ter apenas um motivo para mudar

A class/função/package/componente é especializada em um único assunto.

*/

type BankAccount struct{}

func (b *BankAccount) PrintAccount() {}

func (b *BankAccount) ValidateAccount() {
	//Crucial business logic.
}

func (b *BankAccount) Save() {
	println("Saving data into Database...")
}

/*
Solução:

Devemos separar as responsabilidades, nesse cenário
isolar BankAccount para sua regra de negócio

BankAccount precisa ser responsável apenas do dominio dele.
BankAccount não deve chamar diretamente Save, pois não faz parte do dominio dele

*/

type BankAccountRepository struct{}

func (b *BankAccountRepository) Save() {
	println("Saving data into Database...")
}

type BankAccount2 struct {
	repository BankAccountRepository
}

func (b *BankAccount2) ValidateAccount() {
	//Crucial business logic.
}

func (b *BankAccount2) Save() {
	b.repository.Save()
}

type BankViewer struct{}

func (b *BankViewer) PrintAccount() {}
