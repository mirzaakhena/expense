package repository

import "expense/model"

type Repository interface {
	InsertUser(user *model.User) error
	QueryAllUser(page, size int) ([]*model.User, error)
	InsertExpense(expense *model.Expense) error
	QueryAllExpense(page, size int, userID string) ([]*model.Expense, error)
}
