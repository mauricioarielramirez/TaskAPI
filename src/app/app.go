package app

import (
	"TaskAPI/src/controller"
	"TaskAPI/src/domain"
	"TaskAPI/src/mock"
	"fmt"
)

func DoSomething () {
	//inicializo el mock de cosas
	mock.SetEntities()

	var newUser domain.User
	newUser.SetId(2);
	newUser.SetName("Daniela Jacuzzi")
	newUser.SetAlias("djacuzzi")
	newUser.SetDescription("Service desk employee")

	EmulateOperation("add",newUser)
	fmt.Printf("%+v",controller.ListUsers())
}

func InitialiceRoutes() {

}

func EmulateOperation(operation string, entityUser domain.User) {
	switch operation {
	case "add":
		controller.AddUser(entityUser)
	default:
		fmt.Println("Opción inválida")
	}
}