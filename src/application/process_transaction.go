package application

import "github.com/lukkas-lukkas/go-api-rest/src/domain"

func processTransaction(transaction *domain.Transaction, consumer *domain.Consumer) string {
	error := consumer.Pay(transaction.Amount)

	if error == nil {
		return "APPROVED"
	}

	return "REPROVED"
}
