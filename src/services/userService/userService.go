package userService

import (
	"TaskAPI/src/dao/userDao"
	"TaskAPI/src/domain"
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
	return userDao.List()
}

func DeleteUser(id int) bool {
	return userDao.Delete(id)
}


/*func GetUser (id int) domain.User {
	return nil
}*/