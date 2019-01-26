package controller

import (
	"TaskAPI/src/domain"
	uService "TaskAPI/src/services/userService"
)

func ListUsers() []domain.User {
	return uService.GetUserList()
}