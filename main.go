package main

import (
	"gorm_example/dal"
	"gorm_example/model"
)

func init() {
	dal.Initialize()
}

func main() {
	//createStudent()
}

func createStudent() {
	model.CreateTable()
	model.CreateStudent()
	model.CreateStudentInBatch()
}
