package domain

type Consumer struct {
	ID string
	Balance float64
}

func NewConsumer(id string, balance float64) *Consumer {
	return &Consumer{
		ID: id,
		Balance: balance,
	}
}

func (c *Consumer) Pay(value float64) {
	c.Balance -= value
}
