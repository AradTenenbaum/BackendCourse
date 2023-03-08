package gapi

import (
	"context"
	"database/sql"

	"github.com/AradTenenbaum/BackendCourse/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	account, err := server.store.GetAccount(ctx, int64(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Account not found")
		}
		return nil, status.Errorf(codes.Internal, "Can't get account: %w", err)
	}

	if account.Owner != authPayload.Username {
		return nil, status.Errorf(codes.PermissionDenied, "account doesn't belong to the authenticated user")
	}

	return &pb.GetAccountResponse{Account: convertAccount(account)}, nil
}
