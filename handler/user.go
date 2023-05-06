package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/kikils/vercel-api-sample/domain/model"
	"github.com/kikils/vercel-api-sample/interface/gateway/db"
	"github.com/kikils/vercel-api-sample/openapi"
	"golang.org/x/xerrors"
)

type UserHandler struct {
	db.UserRepository
}

func NewUserHandler(d *sqlx.DB) *UserHandler {
	r := db.NewUserRepository(d)
	return &UserHandler{
		UserRepository: *r,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	var req openapi.PostUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}

	user, err := model.NewUser(req.Name)
	if err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}
	err = h.UserRepository.Store(ctx, user)
	if err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}

	res := openapi.PostUserResponse{
		User: openapi.User{
			Id:   &user.ID,
			Name: user.Name,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}

	return nil
}

func (h *UserHandler) SearchUser(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	user, err := h.UserRepository.SearchByID(ctx, r.URL.Query().Get("id"))
	if err != nil {
		return xerrors.Errorf("UserHandler.SearchUser: %w", err)
	}

	res := openapi.GetUserResponse{
		User: &openapi.User{
			Id:   &user.ID,
			Name: user.Name,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}

	return nil
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	var req openapi.PatchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return xerrors.Errorf("UserHandler.UpdateUser: %w", err)
	}
	id := r.URL.Query().Get("id")

	err := h.UserRepository.UpdateByID(ctx, id, &model.User{Name: req.Name})
	if err != nil {
		return xerrors.Errorf("UserHandler.SearchUser: %w", err)
	}
	updatedUser, err := h.UserRepository.SearchByID(ctx, id)
	if err != nil {
		return xerrors.Errorf("UserHandler.SearchUser: %w", err)
	}

	res := openapi.PatchUserResponse{
		User: &openapi.User{
			Id:   &updatedUser.ID,
			Name: updatedUser.Name,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}

	return nil
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()

	id := r.URL.Query().Get("id")

	err := h.UserRepository.DeleteByID(ctx, id)
	if err != nil {
		return xerrors.Errorf("UserHandler.SearchUser: %w", err)
	}

	res := openapi.DeleteUserResponse{
		User: &openapi.User{
			Id: &id,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		return xerrors.Errorf("UserHandler.CreateUser: %w", err)
	}

	return nil
}
