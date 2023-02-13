package application

import . "github.com/lukkas-lukkas/go-api-rest/src/domain"

func processTransaction(transaction *Transaction, consumer *Consumer) string {
	err := consumer.Pay(transaction.Amount)

	if err == nil {
		return "APPROVED"
	}

	return "REPROVED"
}
