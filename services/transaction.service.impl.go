package services

import (
	"context"
	"database/sql"
	"time"
	"wallet/models"
)

type TransactionServiceImpl struct {
	transactionDB *sql.DB
	ctx context.Context
}

func (t TransactionServiceImpl) CreateTransaction(transaction *models.Transaction) error {
	tx, err := t.transactionDB.BeginTx(t.ctx, nil)
	if err != nil {
		return err
	}
	if _, err := tx.ExecContext(t.ctx, "INSERT INTO transaction (senderName, receiverName, date, money, message, create_at) VALUES (?, ?, ?, ?, ?, ?)", transaction.SenderName, transaction.ReceiverName, transaction.Date, transaction.Money, transaction.Message, time.Now()); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	if _, err := tx.ExecContext(t.ctx, "UPDATE customer SET balance = balance + ?, update_at = ? WHERE ? = (SELECT username FROM user WHERE user.id = customer.id)", transaction.Money, time.Now(), transaction.ReceiverName); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	if _, err := tx.ExecContext(t.ctx, "UPDATE customer SET balance = balance - ?, update_at = ? WHERE ? = (SELECT username FROM user WHERE user.id = customer.id)", transaction.Money, time.Now(), transaction.SenderName); err != nil {
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

func (t TransactionServiceImpl) GetTransactions(customerID string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	rows, errQuery := t.transactionDB.QueryContext(t.ctx, "SELECT id, date, receiverName, money, status, create_at FROM transaction WHERE senderName = (SELECT user.username FROM user, customer WHERE user.id = customer.id AND customer.id = ?)", customerID)
	if errQuery != nil {
		return transactions, errQuery
	}
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Date, &transaction.ReceiverName, &transaction.Money, &transaction.Status, &transaction.CreateAt)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t TransactionServiceImpl) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	rows, errQuery := t.transactionDB.QueryContext(t.ctx, "SELECT * FROM transaction")
	if errQuery != nil {
		return transactions, errQuery
	}
	defer rows.Close()
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.SenderName, &transaction.ReceiverName, &transaction.Date, &transaction.Money, &transaction.Message, &transaction.Status, &transaction.CreateAt)
		if err != nil {
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t TransactionServiceImpl) DeleteTransaction(id string) error {
	_, errExec := t.transactionDB.ExecContext(t.ctx, "DELETE FROM transaction WHERE id = ?", id)
	if errExec != nil {
		return errExec
	}
	return nil
}

func NewTransactionService(transactionDB *sql.DB, ctx context.Context) TransactionService {
	return &TransactionServiceImpl{
		transactionDB: transactionDB,
		ctx: ctx,
	}
}



