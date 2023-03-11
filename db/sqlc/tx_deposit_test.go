package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/AradTenenbaum/BackendCourse/util"
	"github.com/stretchr/testify/require"
)

func TestDepositTx(t *testing.T) {
	store := NewStore(testDB)
	account := createRandomAccount(t)
	currentBalance := account.Balance
	amountToAdd := util.RandomMoney()
	fmt.Printf("account: %s balance: %d toAdd: %d\n", account.Owner, currentBalance, amountToAdd)

	result, err := store.DepositTx(context.Background(), DepositTxParams{
		AccountId: account.ID,
		Amount:    amountToAdd,
	})

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.Account)
	require.NotEmpty(t, result.NewEntry)
	require.Equal(t, result.Account.ID, account.ID)
	require.Equal(t, result.NewEntry.AccountID, account.ID)
	require.Equal(t, result.Account.Balance, currentBalance+amountToAdd)
	require.Equal(t, result.NewEntry.Amount, amountToAdd)

	amountToAdd = util.RandomMoney()

	result, err = store.DepositTx(context.Background(), DepositTxParams{
		AccountId: account.ID,
		Amount:    (-1) * amountToAdd,
	})
	require.Error(t, err)

}
