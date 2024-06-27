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
	readStudentById(2)
	readStudentByStruct()
}

func createStudent() {
	model.CreateTable()
	model.CreateStudent()
	model.CreateStudentInBatch()
}

func readStudentById(studentId uint) {
	model.ReadStudentById(studentId)
}

func readStudentByStruct() {
	model.ReadStudentByStruct()
}
