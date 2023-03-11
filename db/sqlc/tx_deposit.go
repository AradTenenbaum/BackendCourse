package db

import (
	"context"
	"errors"
)

type DepositTxParams struct {
	AccountId int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type DepositTxResult struct {
	NewEntry Entry   `json:"new_entry"`
	Account  Account `json:"account"`
}

func (store *SQLStore) DepositTx(ctx context.Context, arg DepositTxParams) (DepositTxResult, error) {
	var result DepositTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.Amount <= 0 {
			return errors.New("amount must be greater than 0")
		}

		result.NewEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.AccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		result.Account, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Amount: arg.Amount,
			ID:     arg.AccountId,
		})
		return err
	})

	return result, err
}
