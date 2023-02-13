package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	_, error := NewTransaction("1", "1010", 2000)

	assert.Error(t, error)
	assert.Equal(t, "You can't create transaction with amount greater than 1000", error.Error())
}

func TestTransactionWithAmountLessThan1(t *testing.T) {
	_, error := NewTransaction("1", "1010", 0)

	assert.Error(t, error)
	assert.Equal(t, "You can't create transaction with amount less than 1", error.Error())
}