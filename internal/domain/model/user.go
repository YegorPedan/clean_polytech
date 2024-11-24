package model

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	FamilyName string      `json:"family_name"`
	PhoneID    string      `json:"phone_id"`
	Phone      *Smartphone `json:"phone"`
}
