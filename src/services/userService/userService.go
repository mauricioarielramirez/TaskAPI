package userService

import (
	domain "TaskAPI/src/domain"
	"fmt"
)

func AddUser (u domain.User) bool {
	fmt.Println("Elemento agregado")
	fmt.Println(u)
	return true
}

func GetUserList () []domain.User {
	var usuarios [1]domain.User
	var usuario domain.User
	usuario.SetId(1)
	usuario.SetName("Mauricio Ariel Ramirez")
	usuario.SetAlias("marramirez")
	usuario.SetDescription("Usuario con permisos estandar de IT")
	usuarios[0] = usuario
	var listUser = usuarios[0:1]
	return listUser
}