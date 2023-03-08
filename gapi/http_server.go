package gapi

import (
	db "github.com/AradTenenbaum/BackendCourse/db/sqlc"
	"github.com/AradTenenbaum/BackendCourse/util"
)

type HttpServer struct {
	config util.Config
	store  db.Store
}

func NewHttpServer(config util.Config, store db.Store) (*HttpServer, error) {
	return &HttpServer{
		config: config,
		store:  store,
	}, nil
}
