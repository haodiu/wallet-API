package services

import (
	"context"
	"database/sql"
	"time"
	"wallet/models"
)

type UserServiceImpl struct {
	userDB *sql.DB
	ctx context.Context
}

func (u UserServiceImpl) CreateUser(user *models.User) error {
	tx, err := u.userDB.BeginTx(u.ctx, nil)
	if err != nil {
		return err
	}
	result, errExec := tx.ExecContext(u.ctx, "INSERT INTO user(username, email, password, create_at) VALUES (?, ?, ?, ?)", user.Username, user.Email, user.Password, time.Now())
	if errExec != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return errExec
	}
	newCustomerID, errGetLastID := result.LastInsertId()
	if errGetLastID != nil {
		return errGetLastID
	}

	if _, err := tx.ExecContext(u.ctx, "INSERT INTO customer(id, create_at) VALUES(?, ?)", newCustomerID, time.Now()); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	if commitErr := tx.Commit(); commitErr != nil {
		return commitErr
	}
	return nil
}

func (u UserServiceImpl) GetUser(id string) (*models.User, error) {
	var user models.User
	err := u.userDB.QueryRowContext(u.ctx, "SELECT * FROM user WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreateAt, &user.UpdateAt, &user.IsAdmin, &user.IsAddInfo)
	if err == sql.ErrNoRows {
		return &user, err
	}
	return &user, nil
}

func (u UserServiceImpl) CheckUser(user *models.User) (*models.User, error) {
	var result models.User
	errQuery := u.userDB.QueryRowContext(u.ctx, "SELECT id, userName, isAdmin, isAddInfo FROM user WHERE (username = ? OR email = ?) AND password = ?", &user.Username, &user.Email, &user.Password).Scan(&result.ID, &result.Username, &result.IsAdmin, &result.IsAddInfo)
	if errQuery == sql.ErrNoRows {
		return &result, errQuery
	}
	return &result, nil
}

func (u UserServiceImpl) UpdateUser(id string, user *models.User) error {
	_, errExec := u.userDB.ExecContext(u.ctx, "UPDATE user SET username = ?, email = ?, password = ?, isAdmin = ?, update_at = ? WHERE id = ?", &user.Username, &user.Email, &user.Password, &user.IsAdmin, time.Now(), id)
	if errExec != nil {
		return errExec
	}
	return nil
}

func NewUserService(userDB *sql.DB, ctx context.Context) UserService {
	return &UserServiceImpl{
		userDB: userDB,
		ctx: ctx,
	}
}
