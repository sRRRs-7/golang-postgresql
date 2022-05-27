package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/sRRRs-7/golang-postgresql/db/sqlc"
	"github.com/sRRRs-7/golang-postgresql/protobuf"
	"github.com/sRRRs-7/golang-postgresql/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *protobuf.CreateUserRequest) (*protobuf.CreateUserResponse, error) {
	hashedPassword, err := util.HashedPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		Fullname:       req.GetFullname(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &protobuf.CreateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}
