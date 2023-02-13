package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConsumerPayValueGreaterThanHisBalance(t *testing.T) {
	consumer := Consumer{
		ID: "10",
		Balance: 30,
	}

	error := consumer.Pay(31)

	assert.Equal(t, "Value greater than balance for this consumer", error.Error())
}

func TestConsumerPayValueLessThanHisBalance(t *testing.T) {
	consumer := Consumer{
		ID: "10",
		Balance: 30,
	}

	error := consumer.Pay(29)

	assert.Nil(t, error)
}
