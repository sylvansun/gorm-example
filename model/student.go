package model

import (
	"fmt"
	"gorm_example/dal"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func CreateTable() {
	err := dal.DB.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateStudent() {
	email := "jack@gmail.com"
	sampleStu := Student{
		Name:   "jack",
		Age:    18,
		Gender: true,
		Email:  &email,
	}
	err := dal.DB.Create(&sampleStu).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func CreateStudentInBatch() {
	var studentList []Student
	email := "jack@gmail.com"
	for i := 0; i < 100; i++ {
		studentList = append(studentList, Student{
			Name:   fmt.Sprintf("%d", i+1),
			Age:    i,
			Gender: true,
			Email:  &email,
		})
	}
	err := dal.DB.Create(&studentList).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ReadStudentById(studentID uint) {
	var student Student
	err := dal.DB.Take(&student, studentID).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(student)
}

func ReadStudentByStruct() {
	student := Student{}
	student.Name = "jack"
	err := dal.DB.Take(&student).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(student)
}
