package controllers

import (
	"github.com/higayu624/portfolio_go/src/app/interfaces/requests"
	"github.com/higayu624/portfolio_go/src/app/usecase"
	"github.com/gin-gonic/gin"
)

type PlaylistController struct {
	Interactor usecase.PlaylistInteractor
}

func NewPlaylistController() *PlaylistController {
	return &PlaylistController{
		Interactor: usecase.PlaylistInteractor{
			PlaylistRepository: &requests.PlaylistRepository{
			},
		},
	}
}

func (controller *PlaylistController) Index(c *gin.Context) {
	playlist := controller.Interactor.Playlist()
	// if err != nil {
	// 	c.JSON(500, err)
	// 	return//500の場合この関数から抜ける
	// }
	c.JSON(200, playlist)
}