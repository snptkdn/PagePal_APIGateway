package dto

type BookRequestDto struct {
	ISBN   string `json:"ISBN"`
  Title string `json:"title"`
  Author string `json:"author"`
  ImageURL string `json:"image_url"`
  Description string `json:"description"`
  PageCount int `json:"page_count"`
}
