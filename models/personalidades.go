package models

// escrevendo tys e dando enter gera esse código
type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
