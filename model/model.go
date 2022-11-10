package model

import (
	"fmt"
	"strings"
	"time"
)

type Expense struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	UserID      string    `json:"-" gorm:"size:36"`
	User        *User     `json:"user"`
}

func (e Expense) Validate() error {

	if strings.TrimSpace(e.UserID) == "" {
		return fmt.Errorf("user_id must not empty")
	}

	if strings.TrimSpace(e.Description) == "" {
		return fmt.Errorf("description must not empty")
	}

	if e.Amount <= 0 {
		return fmt.Errorf("amount must > 0")
	}

	return nil
}

type User struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (e User) Validate() error {

	if strings.TrimSpace(e.Name) == "" {
		return fmt.Errorf("name must not empty")
	}

	return nil
}
