package models

type Personality struct {
	Id      int    `json:"id"`
	Name    string `json:"name"` // Indicating that this field will be serialized
	History string `json:"history"`
}

var Personalitys []Personality
