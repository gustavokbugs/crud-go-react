package dto

type CreatePerson struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"`
}
