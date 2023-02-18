package entity

type Media struct {
	Description string `json:"description,omitempty"`
	Id          int    `json:"id"`
	Image       string `json:"image"`
}
