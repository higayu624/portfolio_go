package usecase

import (
)

type PlaylistInteractor struct {
	PlaylistRepository PlaylistRepository
}

func (interactor *PlaylistInteractor) Playlist() (playlist []byte) {
	playlist = interactor.PlaylistRepository.GetOne()
	return
}