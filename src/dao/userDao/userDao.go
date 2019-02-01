package userDao

import (
	"TaskAPI/src/dataBaseConnection"
	"TaskAPI/src/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var connection *sqlx.DB


func Add(user domain.User) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustBegin()
	err:=connection.MustExec("INSERT INTO user(id,name,alias,description) VALUES (?, ?, ?, ?)" ,getNextId(),user.Name,user.Alias,user.Description)
	if err!=nil {
		return false
	}
	return true
}

func Modify(user domain.User) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustExec("UPDATE user set name = ?, alias = ?, description = ? where id = ?", user.Name,user.Alias,user.Description,user.Id )
	return true
}

func Delete (id int) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustExec("delete from user where id = ?",id)
	return true
}

func GetById (id int) domain.User {
	var user domain.User
	_,connection = dataBaseConnection.OpenConnection()
	sqlResult,_ := connection.Queryx("SELECT * FROM user where id = ?",id)

	for sqlResult.Next() {
		sqlResult.StructScan(&user)
	}
	return user
}

func List() []domain.User {
	_,connection = dataBaseConnection.OpenConnection()
	users:= []domain.User{}
	err:=connection.Select(&users,"SELECT id, name, alias, description FROM user")
	if err!=nil {
		fmt.Println(err)
	}

	return users
}

func getNextId() int {
	var returnId int
	dataBaseConnection.OpenConnection()
	sqlResult,_ := connection.Queryx("SELECT MAX(ID) as id FROM user")
	for sqlResult.Next() {
		sqlResult.Scan(&returnId)
	}
	return returnId + 1
}

func GetByCriteria (user domain.User) []domain.User {
	_,connection = dataBaseConnection.OpenConnection()

	preparedStatement:="SELECT id, name, alias, description FROM user where "
	users:= []domain.User{}
	params := []string{}

	if user.Name != "" {
		preparedStatement = preparedStatement + "name = ? "
		params = append(params, user.Name)
	} else {
		preparedStatement = preparedStatement + "0 = ?"
		params = append(params, "0")
	}

	if user.Alias!= "" {
		preparedStatement = preparedStatement + " and alias = ? "
		params = append(params, user.Alias)
	} else {
		preparedStatement = preparedStatement + " and 0 = ? "
		params = append(params, "0")
	}

	if user.Description!= "" {
		preparedStatement = preparedStatement + " and description = ? "
		params = append(params, user.Description)
	} else {
		preparedStatement = preparedStatement + " and 0 = ? "
		params = append(params,"0" )
	}

	err := connection.Select(&users,preparedStatement, params[0],params[1],params[2])

	if err!=nil {
		fmt.Println(err)
	}

	return users
}