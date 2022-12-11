package controllers

import (
	"fmt"
	"strconv"

	"github.com/higayu624/portfolio_go/src/app/domain"
	"github.com/higayu624/portfolio_go/src/app/interfaces/requests"
	"github.com/higayu624/portfolio_go/src/app/usecase"
)

type PlaylistController struct {
	Interactor usecase.PlaylistInterator
}

func NewPlaylistController() *PlaylistController {
	return &PlaylistController{
		Interactor: usecase.PlaylistInteractor{
			PlaylistRepository: &requests.PlaylistRepository{
				RequstHandler: RequestHandler
			},
		},
	}
}

func (controller *PlaylistController) Index(c Context) {
	playlist, err := controller.Interactor.Playlist()
	if err != nil {
		c.JSON(500, NewError(err))
		return//500の場合この関数から抜ける
	}
	c.JSON(200, playlist)
}