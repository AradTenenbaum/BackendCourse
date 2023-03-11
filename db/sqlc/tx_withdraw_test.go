package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/AradTenenbaum/BackendCourse/util"
	"github.com/stretchr/testify/require"
)

func TestWithdrawTx(t *testing.T) {
	store := NewStore(testDB)
	account := createRandomAccount(t)
	currentBalance := account.Balance
	amountToDecrease := util.RandomMoney()
	fmt.Printf("account: %s balance: %d toDecrease: %d\n", account.Owner, currentBalance, amountToDecrease)

	result, err := store.WithdrawTx(context.Background(), WithdrawTxParams{
		AccountId: account.ID,
		Amount:    amountToDecrease,
	})

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.Account)
	require.NotEmpty(t, result.NewEntry)
	require.Equal(t, result.Account.ID, account.ID)
	require.Equal(t, result.NewEntry.AccountID, account.ID)
	require.Equal(t, result.Account.Balance, currentBalance-amountToDecrease)
	require.Equal(t, result.NewEntry.Amount, (-1)*amountToDecrease)

	amountToDecrease = util.RandomMoney()

	result, err = store.WithdrawTx(context.Background(), WithdrawTxParams{
		AccountId: account.ID,
		Amount:    (-1) * amountToDecrease,
	})
	require.Error(t, err)

}
