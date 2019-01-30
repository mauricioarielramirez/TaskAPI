package dataBaseConnection

import (
	_ "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

//mapping de entidades


func ConnectToDatabase () {

	var schema=make([]string,3)

	schema = append(schema,"create table user (id int, name varchar(30), alias varchar(30), description varchar(120))",
		"create table task (id int, name varchar(30), description varchar(120), created datetime, modified datetime, modifiedBy int, creator int)",
	"create table user_task(userId int, taskId int)")

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
				db.MustExec(v)
			}

		}
	}
}



	//sqlx.Connect("mysql","user=ariel dbname=goPractice pass=LimaHotel96*")



