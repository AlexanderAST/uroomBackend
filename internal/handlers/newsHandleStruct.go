package handlers

type ReqFindByID struct {
	ID int `json:"id"`
}

type ReqNews struct {
	Date             string `json:"date"`
	Name             string `json:"name"`
	SmallDescription string `json:"small_description"`
	FullDescription  string `json:"full_description"`
	ImagePath        string `json:"image_path"`
}

type ReqUpdateNews struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	Name             string `json:"name"`
	SmallDescription string `json:"small_description"`
	FullDescription  string `json:"full_description"`
	ImagePath        string `json:"image_path"`
}
