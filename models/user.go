package models

type User struct {
	UserName string `json:"username,omitempty" gorm:"primarykey;"`
	Password string `json:"password,omitempty"`
}
