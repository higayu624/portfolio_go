package infrastructure

import (
	"github.com/higayu624/portfolio_go/src/app/domain"
	"github.com/higayu624/portfolio_go/src/app/requests"
	"fmt"
	"net/http"
	"io/ioutil"
)

type RequestHandler struct {}

func (handler *RequestHandler) Request(url string) (body domain.Playlist, err error){
	request, err := http.NewRequest("GET", "https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=RDMMRgKAFK5djSk&key=AIzaSyBL_3aORl9f_AenKMBhEyEuRC83jm7sWYw", nil)
	if err != nil {
		return request, err
	}
  	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.println(body, err)
	return body, err
}