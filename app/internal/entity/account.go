// nome do pacote (está sendo utilizado o nome da referida pasta)
package entity

// dependências
import (
	"time"

	"github.com/google/uuid"
)

// definindo a estrutura (similar à classe)
type Account struct {
	ID        string
	Client    *Client
	ClientID  string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// definindo o método contrutor
// devem ser descritos os argumentos e retornos
func NewAccount(client *Client) *Account {
	// se o client for nulo, não retorna nada
	if client == nil {
		return nil
	}
	// criando
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// retorna a estrutura
	return account
}

// função de crédito
// devem ser descritos a estrutura associada, os argumentos e retornos
func (a *Account) Credit(amount float64) {
	// incrementa o balance
	a.Balance += amount
	// atualiza o updatedAt
	a.UpdatedAt = time.Now()
}

// função de débito
// devem ser descritos a estrutura associada, os argumentos e retornos
func (a *Account) Debit(amount float64) {
	// decrementa o balance
	a.Balance -= amount
	// atualiza o updatedAt
	a.UpdatedAt = time.Now()
}