package app

import (
	"TaskAPI/src/controller"
	"TaskAPI/src/domain"
	"fmt"
)

func DoSomething () {
	//inicializo el mock de cosas
	//mock.SetEntities()
	//dataBaseConnection.CreateSchema()

	var newUser domain.User
	newUser.Id = 0; //con cero voy a indicar que es nuevo
	newUser.Name = "Daniela Jacuzzi"
	newUser.Alias = "djacuzzi"
	newUser.Description = "Service desk employee"

	EmulateOperation("add",newUser)
	fmt.Println("Primera inserción")
	fmt.Printf("%+v",controller.ListUsers())

	var modifiedUser domain.User
	modifiedUser.Id = 2; //voy a llamar a modificar
	modifiedUser.Name = "Eugenia Rubio"
	modifiedUser.Alias = "erubio"
	modifiedUser.Description = "People IT Manager"

	EmulateOperation("modify",modifiedUser)

	fmt.Println("Impresión despues de modificar")
	fmt.Printf("%+v",controller.ListUsers())

	EmulateOperation("delete", domain.User{})
}

func InitialiceRoutes() {

}

func EmulateOperation(operation string, entityUser domain.User) {
	switch operation {
	case "add":
		controller.AddUser(entityUser)
	case "modify":
		controller.AddUser(entityUser)
	case "delete":
		controller.DeleteUser(5)
	case "list":
		controller.ListUsers();
	default:
		fmt.Println("Opción inválida")
	}
}