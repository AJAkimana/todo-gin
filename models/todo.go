package models

type Todo struct {
	ID   int    `json:"id"`
	Tast string `json:"task"`
	Done bool   `json:"done"`
}
