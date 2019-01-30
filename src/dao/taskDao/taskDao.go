package taskDao

import (
	"TaskAPI/src/domain"
	"TaskAPI/src/mock"
)

func Add(task domain.Task) bool {
	mock.TaskList = append(mock.TaskList, task)
	return true
}

func Modify(task domain.Task) bool {
	return true
}

func Delete (id int) bool {
	return true
}

func GetById (id int) domain.Task {
	return domain.Task{}
}

func List() []domain.Task {
	return []domain.Task{}
}
