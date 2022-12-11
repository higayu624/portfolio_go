package usecase

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
)

type PlaylistInteractor struct {
	PlaylistRepository PlaylistRepository
}

func (interactor *PlaylistInteractor) Playlist() (playlist domain.Playlist, err error) {
	playlist, err = interactor.PlaylistRepository.GetOne()
	return
}