package srp

import "fmt"

/*

Single Responsiblity Principle (Princípio da responsabilidade única)

Uma class/função/package/componente, deve ter apenas um motivo para mudar

A class/função/package/componente é especializada em um único assunto.

*/

// BankAccount viola o SRP por conter múltiplas responsabilidades.
type BankAccount struct {
	AccountNumber string
	Balance       float64
}

// PrintAccount imprime os detalhes da conta, uma responsabilidade que não deveria estar em uma classe de domínio.
func (b *BankAccount) PrintAccount() {
	fmt.Printf("Account Number: %s, Balance: $%.2f\n", b.AccountNumber, b.Balance)
}

// ValidateAccount verifica a validade dos dados da conta.
func (b *BankAccount) ValidateAccount() bool {
	return b.AccountNumber != "" && b.Balance >= 0
}

// Save salva os dados da conta no banco de dados, misturando persistência de dados com a lógica de negócios.
func (b *BankAccount) Save() {
	fmt.Println("Saving data to the database...")
}

/*
Solução:

Devemos separar as responsabilidades, nesse cenário
isolar BankAccount para sua regra de negócio

BankAccount precisa ser responsável apenas do dominio dele.
BankAccount não deve chamar diretamente Save, pois não faz parte do dominio dele

*/

// BankAccount apenas contém dados relacionados à conta.
type SBankAccount struct {
	AccountNumber string
	Balance       float64
}

// AccountValidator fornece métodos para validar dados da conta.
type AccountValidator struct{}

func (v *AccountValidator) ValidateAccount(b *BankAccount) bool {
	return b.AccountNumber != "" && b.Balance >= 0
}

// BankAccountRepository responsável pela persistência de dados da conta.
type BankAccountRepository struct{}

func (r *BankAccountRepository) Save(b *SBankAccount) {
	fmt.Println("Saving account data to the database...")
}

// AccountPrinter responsável por imprimir detalhes da conta.
type AccountPrinter struct{}

func (p *AccountPrinter) PrintAccount(b *BankAccount) {
	fmt.Printf("Account Number: %s, Balance: $%.2f\n", b.AccountNumber, b.Balance)
}
