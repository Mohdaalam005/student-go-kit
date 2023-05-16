package repository

import "github.com/volatiletech/null/v8"

type StudentRequest struct {
	Name   null.String `json:"name"`
	Gender null.String `json:"gender"`
	Dob    null.String     `json:"dob"`
}

type StudentResponse struct {
	ID int `json:"id"`
}

type Student struct {
	ID   int    `json:"id"`
	Name   null.String `json:"name"`
	Gender null.String `json:"gender"`
	Dob    null.String `json:"dob"`
}