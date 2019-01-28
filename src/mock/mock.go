package mock

import "TaskAPI/src/domain"

var UserList []domain.User
var TaskList []domain.Task

func SetEntities () {
	UserList = StartMockUsers()
	TaskList = StartMockTasks()
}

func StartMockUsers() []domain.User {
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

func StartMockTasks() []domain.Task {
	var tareas [1]domain.Task
	var tarea domain.Task
	tarea.SetId(1)
	tarea.SetName("Reordenamiento de backlog")
	tarea.SetDescription("Tarea de rutina para ordenar el backlog de acuerdo a priorirades y borrar eliminados.")
	tarea.SetModified("2019-01-27")
	tarea.SetModifiedBy(UserList[0])
	tarea.SetCreated("2019-01-27")
	tarea.SetCreator(UserList[0])
	tarea.SetSharedWith(UserList[0:1])
	var listTasks = tareas[0:1]
	return listTasks
}

