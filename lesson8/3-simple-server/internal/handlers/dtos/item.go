package dtos

type Item struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}
