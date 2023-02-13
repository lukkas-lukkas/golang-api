package domain

import "errors"

type Consumer struct {
	ID      string
	Balance float64
}

func NewConsumer(id string, balance float64) *Consumer {
	return &Consumer{
		ID:      id,
		Balance: balance,
	}
}

func (c *Consumer) Pay(value float64) error {
	if value > c.Balance {
		return errors.New("value greater than balance for this consumer")
	}

	c.Balance -= value
	return nil
}
