package application

import (
	"github.com/lukkas-lukkas/go-api-rest/src/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessTransactionWhenConsumerHasBalance(t *testing.T) {
	transaction, err := domain.NewTransaction("1", "1", 15)
	consumer := domain.NewConsumer("10", 20)
	result := processTransaction(transaction, consumer)

	assert.Equal(t, 5.0, consumer.Balance)
	assert.Equal(t, "APPROVED", result)
	assert.Empty(t, err)
}

func TestProcessTransactionWhenConsumerDontHasBalance(t *testing.T) {
	transaction, err := domain.NewTransaction("1", "1", 25)
	consumer := domain.NewConsumer("10", 20)
	result := processTransaction(transaction, consumer)

	assert.Equal(t, 20.0, consumer.Balance)
	assert.Equal(t, "REPROVED", result)
	assert.Empty(t, err)
}
