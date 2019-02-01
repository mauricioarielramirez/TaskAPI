package taskService

import (
	"TaskAPI/src/dao/taskDao"
	"TaskAPI/src/domain"
	"time"
)

func AddTask (t domain.Task) bool {
	t.Created = (time.Now().Format("2006-01-02 15:04:05"))
	t.Modified = (time.Now().Format("2006-01-02 15:04:05"))
	return taskDao.Add(t)
}

func ModifyTask (task domain.Task) bool {
	task.Modified = time.Now().String()
	return taskDao.Modify(task)
}

func GetTaskList () []domain.Task {
	//fmt.Println(taskDao.List())
	return taskDao.List()
}

func GetTaskById (id int) domain.Task {
	return taskDao.GetById(id)
}

func DeleteTask(id int) bool {
	return taskDao.Delete(id)
}

func GetTaskByCriteria (task domain.Task) []domain.Task{
	return taskDao.GetByCriteria(task)
}