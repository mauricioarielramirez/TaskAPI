package userDao

import (
	"TaskAPI/src/domain"
	"TaskAPI/src/mock"
)

func Add(user domain.User) bool {
	user.SetId(getNextId())
	mock.UserList = append(mock.UserList,user)
	return true
}

func Modify(user domain.User) bool {
	for i,v:= range mock.UserList {
		if v.Id() == user.Id() {
			mock.UserList[i] = user
		}
	}
	return true
}

func Delete (id int) bool {
	return true
}

func GetById (id int) domain.User {
	for _,v:= range mock.UserList {
		if v.Id() == id {
			return v
		}
	}
	return domain.User{}
}

func List() []domain.User {
	return mock.UserList
}

func getNextId() int {
	returnId := 0
	for _,v:= range mock.UserList {
		returnId = v.Id()
	}
	return returnId +1
}