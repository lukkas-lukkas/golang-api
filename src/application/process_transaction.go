package application

import "github.com/lukkas-lukkas/go-api-rest/src/domain"

func processTransaction(transaction *domain.Transaction, consumer *domain.Consumer) {
	consumer.Pay(transaction.Amount)
}