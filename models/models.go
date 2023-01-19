package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name       string       `json:"name"`
	Gender     string       `json:"gender"`
	Dob        string       `json:"dob"`
	Department []Department `json:"department"`
}

type Department struct {
	gorm.Model
	DepartmentName string `json:"departmentName"`
	StudentId      uint   `gorm:"foreignKey:ID" json:"-"`
}
