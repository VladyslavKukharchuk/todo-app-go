package model

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}
