package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"anthonypillot.com/starter-golang/service"
)

type PhotoController struct {
	service service.PhotoService
}

func NewPhotoController(service service.PhotoService) *PhotoController {
	controller := &PhotoController{
		service: service,
	}
	controller.RegisterRoutes()
	return controller
}

func (controller *PhotoController) RegisterRoutes() {
	http.HandleFunc("/photos", controller.GetPhotos)
}

func (controller *PhotoController) GetPhotos(writer http.ResponseWriter, request *http.Request) {
	photos, err := controller.service.GetPhotos()
	if err != nil {
		fmt.Fprintf(writer, "Error: %v", err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(photos)
}
