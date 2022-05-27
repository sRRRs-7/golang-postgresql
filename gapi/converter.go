package gapi

import (
	db "github.com/sRRRs-7/golang-postgresql/db/sqlc"
	"github.com/sRRRs-7/golang-postgresql/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *protobuf.User {
	return &protobuf.User{
		Username:         user.Username,
		FullName:         user.Fullname,
		Email:            user.Email,
		PasswordChangeAt: timestamppb.New(user.PasswordChangeAt),
		CreatedAt:        timestamppb.New(user.CreatedAt),
	}
}
