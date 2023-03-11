package gapi

import (
	"context"

	db "github.com/AradTenenbaum/BackendCourse/db/sqlc"
	"github.com/AradTenenbaum/BackendCourse/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	account, err := server.store.GetAccount(ctx, req.AccountId)
	if account.Owner != authPayload.Username {
		return nil, status.Errorf(codes.PermissionDenied, "not the commited user's account")
	}

	result, err := server.store.WithdrawTx(ctx, db.WithdrawTxParams{
		AccountId: req.AccountId,
		Amount:    req.Amount,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "withdraw failed: %s", err)
	}
	return &pb.WithdrawResponse{
		NewEntry: convertEntry(result.NewEntry),
		Account:  convertAccount(result.Account),
	}, nil
}
