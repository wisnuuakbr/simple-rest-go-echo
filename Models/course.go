package Models

type Course struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// returning rtable name
func (course *Course) TableName() string {
	return "course"
}