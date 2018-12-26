package user

import (
	"context"
	"database/sql"
	"registration-app/helper"
	"registration-app/model"
	uRepo "registration-app/repository"
)

func NewSQLUserRepo(Conn *sql.DB) uRepo.UserRepo {
	return &mysqlUserRepo {
		Conn: Conn,
	}
}

type mysqlUserRepo struct {
	Conn *sql.DB
}

func (m *mysqlUserRepo) RegisterUser(ctx context.Context, u *model.User) (err error) {
	id := helper.GenerateId()

	query := "INSERT User SET id=?, login=?, password=?, email=?, role=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, id, u.Login, u.Password, u.Email, "user")
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}


func (m *mysqlUserRepo) UpdateAdmin(ctx context.Context, u *model.User) (err error) {
	query := "UPDATE User SET role=? WHERE login=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, u.Role, u.Login)
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}
