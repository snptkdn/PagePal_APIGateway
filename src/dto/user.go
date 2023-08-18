package dto

type UserResponseDto struct {
	UserName   string `json:"UserName"`
	TotalBooks int    `json:"TotalBooks"`
	TotalPages int    `json:"TotalPages"`
}
