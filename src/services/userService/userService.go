package userService

import (
	"TaskAPI/src/dao/userDao"
	"TaskAPI/src/domain"
)

func AddUser (u domain.User) bool {
	return userDao.Add(u)
}

func ModifyUser (user domain.User) bool {
	return userDao.Modify(user)
}

func GetUserList () []domain.User {
	return userDao.List()
}

func GetUserById (id int) domain.User {
	return userDao.GetById(id)
}

func DeleteUser(id int) bool {
	return userDao.Delete(id)
}

func GetUserByCriteria (user domain.User) []domain.User{
	return userDao.GetByCriteria(user)
}

//Bussines validations
func verifyAlias (alias string) {

}

