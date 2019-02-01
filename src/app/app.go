package app

import (
	"TaskAPI/src/controller"
	"TaskAPI/src/domain"
	"fmt"
	"github.com/gin-gonic/gin"
)

func DoSomething () {
	//inicializo el mock de cosas
	//mock.SetEntities()
	//dataBaseConnection.CreateSchema()

	/*var newUser domain.User
	newUser.Id = 0; //con cero voy a indicar que es nuevo
	newUser.Name = "Daniela Jacuzzi"
	newUser.Alias = "djacuzzi"
	newUser.Description = "Service desk employee"*/

	//EmulateOperation("add",newUser)
	//fmt.Println("Primera inserción")
	//fmt.Printf("%+v",controller.ListUsers())

	/*var modifiedUser domain.User
	modifiedUser.Id = 2; //voy a llamar a modificar
	modifiedUser.Name = "Eugenia Rubio"
	modifiedUser.Alias = "erubio"
	modifiedUser.Description = "People IT Manager"*/

	//EmulateOperation("modify",modifiedUser)

	//fmt.Println("Impresión del listado")
	//fmt.Printf("%+v",controller.ListUsers(controller.FIND_BY_CRITERIA,domain.User{0,"Eugenia Rubio","erubio",""},0))

	//EmulateOperation("delete", domain.User{})
	//AddNewTask()
	//fmt.Println(controller.ListTask(controller.RETURN_ALL,domain.Task{},0))
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	InitialiceRoutes()

}

func InitialiceRoutes() {
	Router := gin.Default()
	Router.GET("/user/listAll", controller.ListUserAPI)
	Router.Run("localhost:8080")
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
		controller.ListUsers(controller.RETURN_ALL,domain.User{},0)
	default:
		fmt.Println("Opción inválida")
	}
}

func AddNewTask() {
	var task domain.Task
	task.Id = 0
	task.Name = "Limpiar el backlog"
	task.Description = "Quitar del backlog inmediato y apartarlas para validar con el cliente."
	task.Creator=domain.User{}
	task.SharedWith = []domain.User{}

	controller.AddTask(task)
}