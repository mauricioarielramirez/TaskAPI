package controller

import (
	"TaskAPI/src/domain"
	tService "TaskAPI/src/services/taskService"
	uService "TaskAPI/src/services/userService"
	"github.com/gin-gonic/gin"
)

const (
	RETURN_ALL = "returnAll"
	FIND_BY_ID = "findById"
	FIND_BY_CRITERIA = "findByCriteria"
)

func ListUserAPI(c *gin.Context) {
	c.JSON(200,ListUsers(RETURN_ALL,domain.User{},0))
}

func ListUsers(mode string, user domain.User, id int) []domain.User {
	switch mode {
	case RETURN_ALL:
		return uService.GetUserList();
		break
	case FIND_BY_ID:
		return []domain.User{uService.GetUserById(id)}
		break
	case FIND_BY_CRITERIA:
		return uService.GetUserByCriteria(user)
	}
	return []domain.User{}
}

func AddUser(user domain.User) {
	if user.Id == 0 {
		uService.AddUser(user)
	} else {
		modifyUser(user)
	}
}

func modifyUser(user domain.User) {
	uService.ModifyUser(user)
}

func DeleteUser(id int) {
	uService.DeleteUser(id)
}

// Controller para TASK

func AddTask(task domain.Task) {
	if task.Id == 0 {
		tService.AddTask(task)
	} else {
		modifyTask(task)
	}
}
//no me acuerdo el motivo por el cual prepare de esa forma el task
func modifyTask(task domain.Task) {
	tService.ModifyTask(task)
}

func DeleteTask(id int) {
	tService.DeleteTask(id)
}

func ListTask(mode string, task domain.Task, id int) []domain.Task {
	switch mode {
	case RETURN_ALL:
		return tService.GetTaskList();
		break
	case FIND_BY_ID:
		return []domain.Task{tService.GetTaskById(id)}
		break
	case FIND_BY_CRITERIA:
		return tService.GetTaskByCriteria(task)
	}
	return []domain.Task{}
}