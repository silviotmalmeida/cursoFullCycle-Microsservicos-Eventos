// nome do pacote
package main

// dependências
import (
	"database/sql"
	"fmt"

	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/event"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/repository"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_account"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_client"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/usecase/create_transaction"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/web"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/internal/web/webserver"
	"github.com/silviotmalmeida/cursoFullCycle-Microsservicos-Eventos/pkg/events"

	_ "github.com/go-sql-driver/mysql"
)

// função responsável pela criação de um servidor web
// versão sem o gerenciamento da transação com unity of work - uow
func main() {
	// criando a conexão com o bd
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	// em caso de erro
	if err != nil {
		// encerra a conexão
		panic(err)
	}
	// no fim da execução, fecha a conexão
	defer db.Close()

	// criando o gerenciador de eventos
	eventDispatcher := events.NewEventDispatcher()

	// criando o evento de transactionCreated
	transactionCreatedEvent := event.NewTransactionCreatedEvent()

	// criando os repositories
	clientDb := repository.NewClientRepository(db)
	accountDb := repository.NewAccountRepository(db)
	transactionDb := repository.NewTransactionRepository(db)

	// criando os usecases
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)

	// criando o webserver e definindo a porta a ser utilizada
	port := "8080"
	webserver := webserver.NewWebServer(":" + port)

	// criando os handlers dos endpoints
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	// adicionando os handlers ao webserver e configurando os endpoints
	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	// inicializando o webserver
	fmt.Println("Server is running on port", port)
	webserver.Start()
}
