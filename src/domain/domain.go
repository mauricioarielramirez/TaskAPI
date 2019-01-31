package domain

// TYPE User

type User struct {
	id 			int    `db:"id"`
	name 		string `db:"name"`
	alias		string `db:"alias"`
	description string `db:"description"`
}

func (u *User) Description() string {
	return u.description
}

func (u *User) SetDescription(description string) {
	u.description = description
}

func (u *User) Alias() string {
	return u.alias
}

func (u *User) SetAlias(alias string) {
	u.alias = alias
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Id() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

// END OF USER METHODS

// TYPE Task
type Task struct {
	id int
	name string
	description string
	created string
	modified string
	modifiedBy User
	creator User
	sharedWith []User
}

func (t *Task) SharedWith() []User {
	return t.sharedWith
}

func (t *Task) SetSharedWith(sharedWith []User) {
	t.sharedWith = sharedWith
}

func (t *Task) Creator() User {
	return t.creator
}

func (t *Task) SetCreator(creator User) {
	t.creator = creator
}

func (t *Task) ModifiedBy() User {
	return t.modifiedBy
}

func (t *Task) SetModifiedBy(modifiedBy User) {
	t.modifiedBy = modifiedBy
}

func (t *Task) Modified() string {
	return t.modified
}

func (t *Task) SetModified(modified string) {
	t.modified = modified
}

func (t *Task) Created() string {
	return t.created
}

func (t *Task) SetCreated(created string) {
	t.created = created
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) SetName(name string) {
	t.name = name
}

func (t *Task) Id() int {
	return t.id
}

func (t *Task) SetId(id int) {
	t.id = id
}

//END OF TASK METHODS