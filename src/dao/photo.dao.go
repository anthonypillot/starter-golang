package dao

import (
	"log"

	"anthonypillot.com/starter-golang/model"
	"anthonypillot.com/starter-golang/util"
)

type PhotoDAO struct {
	url string
	// Add fields as necessary, e.g., DB connection
}

func NewPhotoDAO(url string) *PhotoDAO {
	return &PhotoDAO{
		url: url,
	}
}

func (dao *PhotoDAO) GetPhotos() ([]model.Photo, error) {
	photos := []model.Photo{}
	_, err := util.Get(dao.url+"/photos", &photos)
	if err != nil {
		log.Fatalf("Error fetching photos: %v", err)
		return nil, err
	}

	return photos, nil
}
