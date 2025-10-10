package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/database"
	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/event"
	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/usecase/create_account"
	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/usecase/create_client"
	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/usecase/create_transaction"
	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/web"
	"github.com/adrianostankewicz/fc-event-driven-architecture/internal/web/webserver"
	"github.com/adrianostankewicz/fc-event-driven-architecture/pkg/events"
	"github.com/adrianostankewicz/fc-event-driven-architecture/pkg/uow"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	//eventDispatcher.Register("TransactionCreated", handler)
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer("3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
