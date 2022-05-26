package models

type Candidato struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" gorm:"unique"`
	Cargo string `json:"cargo"`
}
