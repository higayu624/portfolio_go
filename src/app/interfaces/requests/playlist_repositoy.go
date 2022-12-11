package requests

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
)

type PlaylistRepository struct {
	RequestHandler
}

func (repository *PlaylistRepository) GetOne() (playlist domain.Playlist, err error) {
	body := repository.RequestHandler.Request()
	// if err != nil {
	// 	return
	// }
	return
}
