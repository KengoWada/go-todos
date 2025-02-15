package models

type BaseModel struct {
	ID        int     `json:"id"`
	Version   int     `json:"-"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"-"`
	DeletedAt *string `json:"-"`
}
