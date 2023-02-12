package domain

import "errors"

type Transaction struct {
	ID string
	AccountID string
	Amount float64
}

func NewTransaction(id string, accountID string, amount float64) (*Transaction, error) {
	if (amount > 1000) {
		return nil, errors.New("You can't create transaction with amount greater than 1000")
	}

	return &Transaction{
		ID: id,
		AccountID: accountID,
		Amount: amount,
	}, nil
}
