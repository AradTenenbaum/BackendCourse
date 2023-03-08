package db

import (
	"context"
	"errors"
	"time"
)

type VerifyEmailTxResult struct {
	User User
}

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg SetUsedVerifyEmailParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		verifyEmail, err := q.SetUsedVerifyEmail(ctx, arg)
		if err != nil {
			return err
		}

		if verifyEmail.ExpiredAt.Before(time.Now()) {
			return errors.New("Verification code expired")
		}

		user, err := q.SetEmailVerified(ctx, verifyEmail.Username)
		if err != nil {
			return err
		}
		result.User = user

		return err
	})

	return result, err
}
