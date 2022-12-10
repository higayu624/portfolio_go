package database

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
)

type TaskRepository struct {
	SqlHandler
}

func (repo *TaskRepository) FindAll() (tasks domain.Tasks, err error) {
	rows, err := repo.Query("SELECT id, title FROM tasks")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		
	}
}