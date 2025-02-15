package models

type User struct {
	BaseModel
	Email    string `json:"email"`
	Password string `json:"-"`
	IsActive bool   `json:"-"`
}

type UserProfile struct {
	BaseModel
	Name   string `json:"name"`
	UserID int    `json:"userId"`
	User   *User  `json:"user"`
}
