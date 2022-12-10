package usecase

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
)

type TaskRepository interface {
	FindAll() (domain.Tasks, error)
}