package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Cpf   string `json:"cpf"`
	Nome  string `json:"nome"`
	Curso string `json:"curso"`
}
