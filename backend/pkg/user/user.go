package user

import (
	"context"
	"fmt"

	"github.com/jestradaramos/group-lift/backend/pkg/domain"
	"github.com/jestradaramos/group-lift/backend/pkg/repo"
	"github.com/jestradaramos/group-lift/backend/pkg/repo/bun"
	"github.com/jestradaramos/group-lift/backend/pkg/repo/models"
)

type UserService struct {
	domain.UnimplementedUserServiceServer
	DB repo.Repository
}

func NewUserServiceServer(db *bun.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (u UserService) AddUser(ctx context.Context, req *domain.AddUserRequest) (*domain.AddUserResponse, error) {
	fmt.Println("We made it in here")
	user := &models.User{Username: req.Username, Password: req.Password}
	user, err := u.DB.AddUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	domainUser := fromModelToDomain(user)
	res := &domain.AddUserResponse{User: domainUser}
	return res, nil
}

func fromModelToDomain(user *models.User) *domain.User {
	domainUser := domain.User{
		Username: user.Username,
		Password: user.Password,
	}

	return &domainUser
}
