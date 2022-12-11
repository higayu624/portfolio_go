package requests

import (
)

type PlaylistRepository struct {
	RequestHandler
}

func (repository *PlaylistRepository) GetOne() (playlist []byte) {
	playlist = repository.Request()
	// if err != nil {
	// 	return
	// }
	return
}
