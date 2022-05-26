package models

type Partido struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Sign string `json:"sign"`
	Name string `json:"name" gorm:"unique"`
}
