package controllers

import (
	"fmt"
	"strconv"

	"github.com/higayu624/portfolio_go/src/app/domain"
	"github.com/higayu624/portfolio_go/src/app/interfaces/database"
	"github.com/higayu624/portfolio_go/src/app/usecase"
)

type TaskController struct {
	Interactor usecase.TaskInterator
}

func (controller *TaskController) Index(c Context) {
	tasks, err := controller.Interactor.Tasks()
	if err != nil {
		c.JSON(500, NewError(err))
		return//500の場合この関数から抜ける
	}
	c.JSON(200, tasks)
}