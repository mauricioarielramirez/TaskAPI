package taskDao

import (
	"TaskAPI/src/dataBaseConnection"
	"TaskAPI/src/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var connection *sqlx.DB

func Add(task domain.Task) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustBegin()
	err:=connection.MustExec("INSERT INTO task(id,name,description,created,modified,modifiedBy,creator) VALUES (?, ?, ?, ?, ?, ?, ?)" ,getNextId(),task.Name,task.Description,task.Created,task.Modified,nil,nil)
	if err!=nil {

		return false
	}
	return true
}

func Modify(task domain.Task) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustExec("UPDATE task set name = ? ,description = ?,created = ?,modified = ?,modifiedBy = ?,creator = ?", task.Name,task.Description,task.Created,task.Modified,task.ModifiedBy,task.Creator)
	return true
}

func Delete (id int) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustExec("delete from task where id = ?",id)
	return true
}

func GetById (id int) domain.Task {
	var task domain.Task
	_,connection = dataBaseConnection.OpenConnection()
	sqlResult,_ := connection.Queryx("SELECT * FROM task where id = ?",id)

	for sqlResult.Next() {
		sqlResult.StructScan(&task)
	}
	return task
}

func List() []domain.Task {
	_,connection = dataBaseConnection.OpenConnection()
	tasks:= []domain.Task{}
	err:=connection.Select(&tasks,"SELECT id, name, description, created FROM task")
	if err!=nil {
		fmt.Println(err)
	}
	//fmt.Println(tasks)
	return tasks
}

func getNextId() int {
	var returnId int
	dataBaseConnection.OpenConnection()
	sqlResult,_ := connection.Queryx("SELECT MAX(ID) as id FROM task")
	for sqlResult.Next() {
		sqlResult.Scan(&returnId)
	}
	return returnId + 1
}

/*
Criteria de task sólo va a buscar por propiedades nativas. Para consultas mixtas
o de agregaciones crear otro método.
 */
func GetByCriteria (task domain.Task) []domain.Task {
	_,connection = dataBaseConnection.OpenConnection()

	preparedStatement:="SELECT id, name, alias, description FROM task where "
	tasks:= []domain.Task{}
	params := []string{}

	if task.Name != "" {
		preparedStatement = preparedStatement + "name = ? "
		params = append(params, task.Name)
		//isFirstCondition = false
	} else {
		preparedStatement = preparedStatement + "0 = ?"
		params = append(params, "0")
	}

	if task.Description!= "" {
		preparedStatement = preparedStatement + " and description = ? "
		params = append(params, task.Description)
	} else {
		preparedStatement = preparedStatement + " and 0 = ? "
		params = append(params, "0")
	}

	if task.Created!= "" {
		preparedStatement = preparedStatement + " and created = ? "
		params = append(params, task.Created)
	} else {
		preparedStatement = preparedStatement + " and 0 = ? "
		params = append(params,"0" )
	}

	if task.Modified!= "" {
		preparedStatement = preparedStatement + " and modified = ? "
		params = append(params, task.Modified)
	} else {
		preparedStatement = preparedStatement + " and 0 = ? "
		params = append(params,"0" )
	}

	err := connection.Select(&tasks,preparedStatement, params[0],params[1],params[2])

	if err!=nil {
		fmt.Println(err)
	}

	return tasks
}