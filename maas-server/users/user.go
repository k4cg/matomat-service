package users

//Represenation of a Matomat users properties, currently only an ID and the Credits (in currency cents).
type User struct {
	ID       uint32
	Username string
	Password string
	Credits  int32
	admin    bool
}

func (u *User) IsAdmin() bool {
	return u.admin
}
