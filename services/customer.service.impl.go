package services

import (
	"context"
	"database/sql"
	"time"
	"wallet/models"
)

type CustomerServiceImpl struct {
	customerDB *sql.DB
	ctx context.Context
}

func (c CustomerServiceImpl) CreateCustomer(customer *models.Customer) error {
	tx, err := c.customerDB.BeginTx(c.ctx, nil)
	if err != nil {
		return err
	}
	result, errAddNewUser := tx.ExecContext(c.ctx, "INSERT INTO user (isAddInfo, create_at) VALUES (true, ?)", time.Now())
	if errAddNewUser != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return errAddNewUser
	}
	lastIdUser, errGetLastID := result.LastInsertId()
	if errGetLastID != nil {
		return nil
	}
	if _, errAddNewCus := tx.ExecContext(c.ctx, "INSERT INTO customer (id, firstName, lastName, dateOfBirth, nationality, address, create_at) VALUES (?, ?, ?, ?, ?, ?, ?)", lastIdUser, customer.FirstName, customer.LastName, customer.DateOfBirth, customer.Nationality, customer.Address, time.Now()); errAddNewCus != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return errAddNewCus
	}
	if commitErr := tx.Commit(); commitErr != nil {
		return commitErr
	}
	return nil
}

func (c CustomerServiceImpl) GetCustomer(id string) (*models.Customer, error) {
	var customer models.Customer
	row := c.customerDB.QueryRowContext(c.ctx, "SELECT * FROM customer WHERE id = ?", id)
	err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.DateOfBirth, &customer.Nationality, &customer.Address, &customer.Balance, &customer.Avatar, &customer.CreateAt, &customer.UpdateAt)
	if err != nil {
		return &customer, err
	}
	return &customer, nil
}

func (c CustomerServiceImpl) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	rows, errQuery := c.customerDB.QueryContext(c.ctx, "SELECT * FROM customer")
	if errQuery != nil {
		return customers, errQuery
	}
	defer rows.Close()
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.DateOfBirth, &customer.Nationality, &customer.Address, &customer.Balance, &customer.Avatar, &customer.CreateAt, &customer.UpdateAt)
		if err != nil {
			return customers, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c CustomerServiceImpl) UpdateCustomer(id string, customer *models.Customer) error {
	_, errExec := c.customerDB.ExecContext(c.ctx, "UPDATE customer SET firstName = ?, lastName = ?, dateOfBirth = ?, nationality = ?,  address= ?, avatar = ?, update_at = ? WHERE id = ?", customer.FirstName, customer.LastName, customer.DateOfBirth, customer.Nationality, customer.Address, customer.Avatar, time.Now(), id)
	if errExec != nil {
		return errExec
	}
	return nil
}

func (c CustomerServiceImpl) DeleteCustomer(id string) error {
	_, errExec := c.customerDB.ExecContext(c.ctx, "DELETE FROM customer WHERE id = ?", id)
	if errExec != nil {
		return errExec
	}
	return nil
}

func NewCustomerService(customerDB *sql.DB, ctx context.Context) CustomerService {
	return &CustomerServiceImpl{
		customerDB: customerDB,
		ctx: ctx,
	}
}

