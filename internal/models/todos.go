package models

type Todo struct {
	BaseModel
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Done        bool         `json:"done"`
	Tags        []*Tags      `json:"tags"`
	OwnerID     int          `json:"ownerId"`
	Owner       *UserProfile `json:"owner"`
}

type Tags struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
}
