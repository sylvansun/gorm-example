package model

import "fmt"

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func CreateStudent() *Student {
	email := "jack@gmail.com"
	sampleStu := Student{
		Name:   "jack",
		Age:    18,
		Gender: true,
		Email:  &email,
	}
	return &sampleStu
}

func CreateStudentInBatch() []Student {
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
	return studentList
}
