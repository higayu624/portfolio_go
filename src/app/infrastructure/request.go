package infrastructure

import (
	"net/http"
	"io"
)

type RequestHandler struct {}

func (handler *RequestHandler) Request() (content []byte, err error){
	request, _ := http.NewRequest("GET", "https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=RDMMRgKAFK5djSk&key=AIzaSyBL_3aORl9f_AenKMBhEyEuRC83jm7sWYw", nil)
	// if err != nil {
	// 	return request, err
	// }
  	client := new(http.Client)
	response, _ := client.Do(request)
	// if err != nil {
	// 	return response, err
	// }
	defer response.Body.Close()
	content, err = io.ReadAll(response.Body)
	return
}