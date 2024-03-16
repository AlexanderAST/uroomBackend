package model

type News struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	Name             string `json:"name"`
	SmallDescription string `json:"small_description"`
	FullDescription  string `json:"full_description"`
	ImagePath        string `json:"image_path"`
}
