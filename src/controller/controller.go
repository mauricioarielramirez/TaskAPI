package controller

import (
	"TaskAPI/src/domain"
	uService "TaskAPI/src/services/userService"
)

func ListUsers() []domain.User {
	return uService.GetUserList()
}

func AddUser(user domain.User) {
	if user.Id() == 0 {
		uService.AddUser(user)
	} else {
		modifyUser(user)
	}
}

func modifyUser(user domain.User) {
	uService.ModifyUser(user)
}