package dto

type UserResponseDto struct {
	UserID     string `json:"UserID"`
	UserName   string `json:"UserName"`
	TotalBooks int    `json:"TotalBooks"`
	TotalPages int    `json:"TotalPages"`
}
