package db

import (
	"context"
	"testing"
	"time"

	"github.com/grayjunzi/backend-master-class-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTrasnferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTrasnfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, account1.ID)
	require.Equal(t, arg.ToAccountID, account2.ID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func createRandomTransfers(t *testing.T) (int64, int64) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTrasnferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	for i := 0; i < 10; i++ {
		transfer, err := testQueries.CreateTrasnfer(context.Background(), arg)

		require.NoError(t, err)
		require.NotEmpty(t, transfer)

		require.Equal(t, arg.FromAccountID, account1.ID)
		require.Equal(t, arg.ToAccountID, account2.ID)
		require.Equal(t, arg.Amount, transfer.Amount)
	}

	return account1.ID, account2.ID
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {

	accountID1, accountID2 := createRandomTransfers(t)

	arg := ListTransfersParams{
		FromAccountID: accountID1,
		ToAccountID:   accountID2,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
