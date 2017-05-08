package templates

type State struct {
	Login *User
}

type User struct {
	Email        string
	Password     string
	PasswordHash string
}
