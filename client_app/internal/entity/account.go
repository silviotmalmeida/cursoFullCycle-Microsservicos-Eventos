// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (

)

// definindo a estrutura (similar à classe)
type Account struct {
	// definindo os atributos e seus tipos
	ID        string
	Balance   float64
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewAccount(id string, balance float64) *Account {
	// criando
	account := &Account{
		ID:        id,
		Balance:   balance,
	}
	// retorna a estrutura
	return account
}

// função de atualização de balance
// devem ser descritos a estrutura associada, os argumentos e retornos
func (a *Account) SetBalance(balance float64) {
	// incrementa o balance
	a.Balance = balance
}
