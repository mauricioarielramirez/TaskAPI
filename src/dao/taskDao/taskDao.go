package taskDao

import (
	"TaskAPI/src/dataBaseConnection"
	"TaskAPI/src/domain"
	"github.com/jmoiron/sqlx"
)

var connection *sqlx.DB

func Add(task domain.Task) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.Begin();
	err:=connection.MustExec("INSERT INTO TASK(id,name,description,created,modified,modifiedBy,creator) values $1,$2,$3,$4,$5,$6,$7",task.Id(),task.Name(),task.Description(),task.Created(),task.Modified(),task.ModifiedBy(),task.Creator())
	if err!=nil {
		return false
	}
	return true
}

func Modify(task domain.Task) bool {
	return true
}

func Delete (id int) bool {
	return true
}

func GetById (id int) domain.Task {
	return domain.Task{}
}

func List() []domain.Task {
	return []domain.Task{}
}
