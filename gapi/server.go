package gapi

import (
	"fmt"

	db "github.com/sRRRs-7/golang-postgresql/db/sqlc"
	"github.com/sRRRs-7/golang-postgresql/protobuf"
	"github.com/sRRRs-7/golang-postgresql/token"
	"github.com/sRRRs-7/golang-postgresql/util"
)

// Server services gRPC routers
type Server struct {
	protobuf.UnimplementedBankAppServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// PASETO -> JWT : NewPasetoMaker -> NewJWTMaker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
