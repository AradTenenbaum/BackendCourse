package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	db "github.com/AradTenenbaum/BackendCourse/db/sqlc"
	"github.com/AradTenenbaum/BackendCourse/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Transfer(ctx context.Context, req *pb.TransferRequest) (*pb.TransferResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	fromAccount, err, code := server.validAccount(ctx, req.FromAccountId, req.Currency)
	if err != nil {
		return nil, status.Errorf(code, "Not a valid account [%d]: %s", req.FromAccountId, err)
	}

	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user account")
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}

	_, err, code = server.validAccount(ctx, req.ToAccountId, req.Currency)
	if err != nil {
		return nil, status.Errorf(code, "Not a valid account [%d]: %s", req.ToAccountId, err)
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID:   req.ToAccountId,
		Amount:        req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Transfer error: %s", err)
	}

	return &pb.TransferResponse{
		Transfer:    convertTransfer(result.Transfer),
		FromAccount: convertAccount(result.FromAccount),
		ToAccount:   convertAccount(result.ToAccount),
		FromEntry:   convertEntry(result.FromEntry),
		ToEntry:     convertEntry(result.ToEntry),
	}, nil
}

func (server *Server) validAccount(ctx context.Context, accountID int64, currency string) (db.Account, error, codes.Code) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, err, codes.NotFound
		}
		return account, err, codes.Internal
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		return account, err, codes.InvalidArgument
	}

	return account, nil, codes.OK
}
