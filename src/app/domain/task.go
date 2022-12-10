package domain

import (
	"time"
)

type Task struct {
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Status		int		`json:"status"`
	Description	string	`json:"description"`
	Deadline	time.Time	`json:"deadline"`
	Startline	time.Time	`json:"startline"`
	Created_at	time.Time	`json:"created_at"`
	Updated_at	time.Time	`json:"updated_at"`
}

type Tasks []Task