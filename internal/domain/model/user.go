package model

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
}
