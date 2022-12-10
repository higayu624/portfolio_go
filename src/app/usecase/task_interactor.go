package usecase

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
	"fmt"
)

type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (interactor *TaskInteractor) Tasks() (task domain.Tasks, err error) {
	task, err = interactor.TaskRepository.FindAll()
	fmt.Println("ここまできてる")
	return
}