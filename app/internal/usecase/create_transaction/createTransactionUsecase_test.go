// nome do pacote (está sendo utilizado o nome da referida pasta)
package create_transaction

// dependências
import (
	"testing"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/entity"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// testes de unidade

// teste de execução com sucesso
func TestCreateTransactionUseCase_Execute(t *testing.T) {
	// criando o client 1
	client1, _ := entity.NewClient("client1", "j@j.com")
	// criando o account 1
	account1 := entity.NewAccount(client1)
	// creditando no account 1
	account1.Credit(1000)
	// criando o client 2
	client2, _ := entity.NewClient("client2", "j@j2.com")
	// criando o account 2
	account2 := entity.NewAccount(client2)
	// creditando no account 2
	account2.Credit(1000)
	// criando o accountMock
	accountGatewayMock := &mocks.AccountGatewayMock{}
	// definindo o retorno do médoto FindByID para a account 1, como account 1 e erro null
	accountGatewayMock.On("FindByID", account1.ID).Return(account1, nil)
	// definindo o retorno do médoto FindByID para a account 2, como account 2 e erro null
	accountGatewayMock.On("FindByID", account2.ID).Return(account2, nil)
	// criando o transactionMock
	transactionGatewayMock := &mocks.TransactionGatewayMock{}
	// definindo o retorno do médoto Save como null
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)
	// criando o usecase
	uc := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)
	// definindo o input
	input := &CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}
	// executando o usecase
	output, err := uc.Execute(input)
	// não deve retornar erro
	assert.Nil(t, err)
	// deve retornar um output válido
	assert.NotNil(t, output)
	// o ID do do output deve estar preenchido
	assert.NotEmpty(t, output.ID)
	// o AccountIDFrom do do output deve corresponder ao input
	assert.Equal(t, input.AccountIDFrom, output.AccountIDFrom)
	// o AccountIDTo do do output deve corresponder ao input
	assert.Equal(t, input.AccountIDTo, output.AccountIDTo)
	// o Amount do do output deve corresponder ao input
	assert.Equal(t, input.Amount, output.Amount)
	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)
	// o método FindByID do mock deve ter sido chamado 2 vezes
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	// o método Create do mock deve ter sido chamado 1 vez
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)

	// mockUow := &mocks.UowMock{}
	// mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	// inputDto := CreateTransactionInputDTO{
	// 	AccountIDFrom: account1.ID,
	// 	AccountIDTo:   account2.ID,
	// 	Amount:        100,
	// }

	// dispatcher := events.NewEventDispatcher()
	// eventTransaction := event.NewTransactionCreated()
	// eventBalance := event.NewBalanceUpdated()
	// ctx := context.Background()

	// uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)
	// output, err := uc.Execute(ctx, inputDto)
	// assert.Nil(t, err)
	// assert.NotNil(t, output)
	// mockUow.AssertExpectations(t)
	// mockUow.AssertNumberOfCalls(t, "Do", 1)
}