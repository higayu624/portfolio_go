package usecase

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
)

type PlaylistRepository interface {
	GetOne() (domain.Playlist, error)
}