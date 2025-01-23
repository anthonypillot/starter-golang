package service

import (
	"anthonypillot.com/starter-golang/dao"
	"anthonypillot.com/starter-golang/model"
)

type PhotoService struct {
	dao dao.PhotoDAO
}

func NewPhotoService(dao dao.PhotoDAO) *PhotoService {
	return &PhotoService{
		dao: dao,
	}
}

func (service *PhotoService) GetPhotos() ([]model.Photo, error) {
	photos, err := service.dao.GetPhotos()
	if err != nil {
		return nil, err
	}

	return photos, err
}
