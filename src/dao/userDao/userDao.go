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
	fmt.Println(connection)
	connection.MustBegin()
	err:=connection.MustExec("INSERT INTO user(id,name,alias,description) VALUES (?, ?, ?, ?)" ,getNextId(),user.Name(),user.Alias(),user.Description())
	if err!=nil {
		return false
	}
	return true
}

func Modify(user domain.User) bool {
	_,connection = dataBaseConnection.OpenConnection()
	connection.MustExec("UPDATE user set name = ?, alias = ?, description = ? where id = ?", user.Name(),user.Alias(),user.Description(),user.Id() )
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
	users:= []*domain.User{}
	//var user domain.User
	//var alias string

	/*sqlResult,_ := connection.Queryx("select id, name, alias, description FROM user")
	for sqlResult.Next() {
		err:=sqlResult.StructScan(&user)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(user)
	}*/
	err:=connection.Get(&users,"SELECT id, name, alias, description FROM user")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(users)

	return []domain.User{}
}

func getNextId() int {
	var returnId int
	dataBaseConnection.OpenConnection()
	sqlResult,_ := connection.Queryx("SELECT MAX(ID) as id FROM user")
	for sqlResult.Next() {
		sqlResult.Scan(&returnId)
	}
	fmt.Println(returnId)
	return returnId + 1
}