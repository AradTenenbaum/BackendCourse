package gapi

import (
	"fmt"

	db "github.com/AradTenenbaum/BackendCourse/db/sqlc"
	"github.com/AradTenenbaum/BackendCourse/pb"
	"github.com/AradTenenbaum/BackendCourse/token"
	"github.com/AradTenenbaum/BackendCourse/util"
	"github.com/AradTenenbaum/BackendCourse/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// Returns a new gRPC server
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
