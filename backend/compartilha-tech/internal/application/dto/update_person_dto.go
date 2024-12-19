package dto

type UpdatePerson struct {
	Name   *string `json:"name"`
	Age    *int    `json:"age"`
	Active *bool   `json:"active"`
}
