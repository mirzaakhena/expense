package controller

import (
	"expense/model"
	"expense/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Data         any    `json:"data"`
	ErrorMessage string `json:"error_message"`
}

type Controller struct {
	Repo repository.Repository
}

func (m Controller) RegisterUser(c *gin.Context) {

	type UserRequest struct {
		Name string `json:"name"`
	}

	var req UserRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	user := model.User{
		ID:   uuid.New().String(),
		Name: req.Name,
	}

	err = user.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	err = m.Repo.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: user, ErrorMessage: ""})
}

func (m Controller) ViewAllUser(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	sizeStr := c.DefaultQuery("size", "3")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	users, err := m.Repo.QueryAllUser(page, size)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: users, ErrorMessage: ""})

}

func (m Controller) AddExpense(c *gin.Context) {

	type ExpenseRequest struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
		UserID      string  `json:"user_id"`
	}

	var req ExpenseRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	expense := model.Expense{
		ID:          uuid.New().String(),
		Amount:      req.Amount,
		Description: req.Description,
		Date:        time.Now(),
		UserID:      req.UserID,
		User:        nil,
	}

	err = expense.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	err = m.Repo.InsertExpense(&expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: expense, ErrorMessage: ""})

}

func (m Controller) ViewAllExpense(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	sizeStr := c.DefaultQuery("size", "3")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	userID := c.DefaultQuery("user_id", "")

	expenses, err := m.Repo.QueryAllExpense(page, size, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: nil, ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{Data: expenses, ErrorMessage: ""})

}
