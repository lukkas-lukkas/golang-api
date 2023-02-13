package application

import (
	"testing"
	"github.com/lukkas-lukkas/go-api-rest/src/domain"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenConsumerHasBalance(t *testing.T) {
	transaction, error := domain.NewTransaction("1", "1", 15)
	consumer := domain.NewConsumer("10", 20)
	processTransaction(transaction, consumer)

	assert.Equal(t, 5.0, consumer.Balance)
	assert.Empty(t, error)
}
