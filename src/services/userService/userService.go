package userService

import (
	domain "TaskAPI/src/domain"
	"TaskAPI/src/mock"
	"fmt"
)

//Slice de users
var userList []domain.User

func AddUser (u domain.User) bool {
	//Agrego elementos en el mock
	mock.UserList = append(mock.UserList,u)
	fmt.Println("Elemento agregado")
	return true
}

func ModifyUser (user domain.User) {

}

func GetUserList () []domain.User {
	//uso las cosas con el mock
	var listUser = mock.UserList
	return listUser
}

/*func GetUser (id int) domain.User {
	return nil
}*/