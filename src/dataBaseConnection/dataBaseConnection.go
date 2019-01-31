package dataBaseConnection

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

//mapping de entidades


func CreateSchema() {

	var schema= append(make([]string, 3), "select * from task")

	schema = append(schema,"create table if not exists user (id int, name varchar(30), alias varchar(30), description varchar(120))",
		"create table if not exists  task (id int, name varchar(30), description varchar(120), created datetime, modified datetime, modifiedBy int, creator int)",
		"create table if not exists  user_task(userId int, taskId int)")

	/*schema = "create table user (id int, name varchar(30), alias varchar(30), description varchar(120)); " +
		"create table task (id int, name varchar(30), description varchar(120), created datetime, modified datetime, modifiedBy int, creator int);" +
		"create table user_task(userId int, taskId int);"*/

	//<username>:<password>@tcp(<host:port>)/<schemaname>
	db,err := sqlx.Connect("mysql","ariel:LimaHotel96*@tcp(localhost:3306)/goPractice")

	if err != nil {
		fmt.Println("Ocurrió un error al intentar conectarse con la base de datos")
		fmt.Println(err)
	} else {
		fmt.Println("Conexión exitosa")
		for _,v:= range schema {
			if v != "" {
				fmt.Println(db.MustExec(v).RowsAffected())
			}

		}
	}
}

func OpenConnection() (string,*sqlx.DB) {
	connection,err := sqlx.Connect("mysql","ariel:LimaHotel96*@tcp(localhost:3306)/goPractice")
	fmt.Println(connection)
	if err!=nil {
		return "Falló la conexión", connection
	} else {
		return "Conexión exitosa", connection
	}
}


