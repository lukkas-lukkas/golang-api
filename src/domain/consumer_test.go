package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumerPayValueGreaterThanHisBalance(t *testing.T) {
	consumer := Consumer{
		ID:      "10",
		Balance: 30,
	}

	err := consumer.Pay(31)

	assert.Equal(t, "value greater than balance for this consumer", err.Error())
}

func TestConsumerPayValueLessThanHisBalance(t *testing.T) {
	consumer := Consumer{
		ID:      "10",
		Balance: 30,
	}

	err := consumer.Pay(29)

	assert.Nil(t, err)
}
