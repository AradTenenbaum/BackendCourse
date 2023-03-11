package db

import (
	"context"
	"errors"
)

type WithdrawTxParams struct {
	AccountId int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type WithdrawTxResult struct {
	NewEntry Entry   `json:"new_entry"`
	Account  Account `json:"account"`
}

func (store *SQLStore) WithdrawTx(ctx context.Context, arg WithdrawTxParams) (WithdrawTxResult, error) {
	var result WithdrawTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.Amount <= 0 {
			return errors.New("amount must be greater than 0")
		}
		fixedAmount := (-1) * arg.Amount

		result.NewEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.AccountId,
			Amount:    fixedAmount,
		})
		if err != nil {
			return err
		}

		result.Account, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			Amount: fixedAmount,
			ID:     arg.AccountId,
		})
		return err
	})

	return result, err
}
