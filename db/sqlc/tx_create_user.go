package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}
type CreateUserTxResult struct {
	User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return nil
		}

		return arg.AfterCreate(result.User)
	})

	return result, err
}
