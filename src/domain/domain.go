package domain
// Es importante que el struct tenga campos en mayúsculas para que encuentre el identificador dónde se exporta
// No se puede utilizar en este casos getter and setters, ya que se accede como propiedad del elemento

// TYPE User
type User struct {
	Id 			int    `db:"id"`
	Name 		string `db:"name"`
	Alias		string `db:"alias"`
	Description string `db:"description"`
}

// END OF USER METHODS

// TYPE Task
type Task struct {
	Id 			int 	`db:"id"`
	Name 		string	`db:"name"`
	Description string	`db:"description"`
	Created 	string	`db:"created"`
	Modified 	string	`db:"modified"`
	ModifiedBy 	User	`db:"modifiedBy"`
	Creator 	User	`db:"creator"`
	SharedWith []User	`db:"sharedWith"`
}


//END OF TASK METHODS