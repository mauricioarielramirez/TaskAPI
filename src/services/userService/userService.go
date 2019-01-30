package userService

import (
	"TaskAPI/src/dao/userDao"
	"TaskAPI/src/domain"
	"TaskAPI/src/mock"
)

//Slice de users
var userList []domain.User

func AddUser (u domain.User) bool {
	return userDao.Add(u)
}

func ModifyUser (user domain.User) bool {
	return userDao.Modify(user)
}

func GetUserList () []domain.User {
	//uso las cosas con el mock
	var listUser = mock.UserList
	return listUser
}

/*func GetUser (id int) domain.User {
	return nil
}*/