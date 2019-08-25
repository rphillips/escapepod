package models

type User struct {
	BaseModel
	Email    string `json:"email" gorm:"not_null;unique_index"`
	Name     string `json:"name" gorm:"not_null"`
	Password string `json:"password,omitempty" gorm:"not null"`
}
