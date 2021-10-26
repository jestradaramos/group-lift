package user

import (
	"context"

	"github.com/jestradaramos/group-lift/backend/pkg/domain"
)

type UserService struct {
	domain.UnimplementedUserServiceServer
}

func (u UserService) AddUser(context.Context, *domain.AddUserRequest) (*domain.AddUserResponse, error) {
	return nil, nil
}
