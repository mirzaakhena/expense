package repository

import (
	"expense/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type GormImpl struct {
	DB *gorm.DB
}

func NewGormImpl() (Repository, error) {
	dsn := "root:mypass123@tcp(127.0.0.1:3306)/expense_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(model.User{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(model.Expense{}); err != nil {
		return nil, err
	}

	return &GormImpl{
		DB: db,
	}, nil
}

func (g GormImpl) InsertUser(user *model.User) error {
	err := g.DB.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (g GormImpl) QueryAllUser(page, size int) ([]*model.User, error) {

	users := make([]*model.User, 0)
	err := g.DB.Limit(size).Offset((page - 1) * size).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (g GormImpl) InsertExpense(expense *model.Expense) error {

	err := g.DB.Save(&expense).Error
	if err != nil {
		return err
	}

	return nil
}

func (g GormImpl) QueryAllExpense(page, size int, userID string) ([]*model.Expense, error) {

	expenses := make([]*model.Expense, 0)

	tx := g.DB.Limit(size).Offset((page - 1) * size).Preload("User")

	if strings.TrimSpace(userID) != "" {
		tx = tx.Where("user_id = ?", userID)
	}

	err := tx.Find(&expenses).Error
	if err != nil {
		return nil, err
	}

	return expenses, nil
}
