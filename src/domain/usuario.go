package domain

import "strings"

type Usuario struct {
	nombre string
	mail string
	nick string
	password string
}

func (user *Usuario) NameIsEmpty() bool{
	return strings.TrimSpace(user.nombre) == ""
}

func (user *Usuario) MailIsEmpty() bool{
	return strings.TrimSpace(user.mail) == ""
}


func (user *Usuario) PasswordIsEmpty() bool{
	return strings.TrimSpace(user.password) == ""
}

func (user *Usuario) AuthUser(nick string,password string )bool{
	return user.nick == nick && user.password == password
}
