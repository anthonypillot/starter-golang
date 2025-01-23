package model

type Photo struct {
	AlbumID      int64  `json:"albumId"`
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
