package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.61

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/paundraP/be-mcs/user-service/graphql/generated"
	"github.com/paundraP/be-mcs/user-service/services"
	"github.com/paundraP/be-mcs/user-service/tools"
	"gorm.io/gorm"
)

// Login is the resolver for the login field.
func (r *authOpsResolver) Login(ctx context.Context, obj *generated.AuthOps, email string, password string) (any, error) {
	getUser, err := r.UserRepo.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("not found")
		}

	}
	isValid := tools.CompareHashPassword(password, getUser.Password)
	if isValid != nil {
		return nil, errors.New("wrong password")
	}

	token, err := services.GenerateJWT(ctx, getUser.ID)
	if err != nil {
		return nil, errors.New("jwt error")
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(ctx context.Context) (*generated.AuthOps, error) {
	return &generated.AuthOps{}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input generated.CreateUserInput) (*generated.User, error) {
	userID := uuid.New()

	hashPassword := tools.HashPassword(input.Password)

	exist := r.UserRepo.CheckEmail(input.Email)
	if exist {
		return nil, errors.New("email already exist")
	}

	user := &generated.User{
		ID:       userID.String(),
		Name:     input.Name,
		Email:    input.Email,
		Password: hashPassword,
	}

	if err := r.UserRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*bool, error) {
	err := r.UserRepo.DeleteUser(id)
	if err != nil {
		failure := false
		return &failure, err
	}
	success := true
	return &success, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*generated.User, error) {
	user, err := r.UserRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &generated.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*generated.GetUserResponse, error) {
	users, err := r.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var res []*generated.GetUserResponse
	for _, user := range users {
		res = append(res, &generated.GetUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return res, nil
}

func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "succes", nil
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
