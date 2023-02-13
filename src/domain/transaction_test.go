package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	_, err := NewTransaction("1", "1010", 2000)

	assert.Error(t, err)
	assert.Equal(t, "you can't create transaction with amount greater than 1000", err.Error())
}

func TestTransactionWithAmountLessThan1(t *testing.T) {
	_, err := NewTransaction("1", "1010", 0)

	assert.Error(t, err)
	assert.Equal(t, "you can't create transaction with amount less than 1", err.Error())
}
