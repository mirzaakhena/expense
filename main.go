package main

import (
	"expense/controller"
	"expense/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	repo, err := repository.NewGormImpl()
	if err != nil {
		panic(err)
	}

	ctrl := controller.Controller{
		Repo: repo,
	}

	router := gin.Default()

	router.POST("/user", ctrl.RegisterUser)
	router.GET("/user", ctrl.ViewAllUser)
	router.POST("/expense", ctrl.AddExpense)
	router.GET("/expense", ctrl.ViewAllExpense)

	err = router.Run()
	if err != nil {
		panic(err)
	}
}
