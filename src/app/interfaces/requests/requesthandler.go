package requests

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
)

type RequestHandler interface {
	Request()(body domain.Playlist)
}