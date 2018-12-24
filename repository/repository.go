package repository

import (
"context"
"registration-app/model"
)

type UserRepo interface {
	RegisterUser(ctx context.Context, u *model.User) (err error)
	UpdateAdmin(ctx context.Context, u *model.User) (user *model.User, err error)
}
