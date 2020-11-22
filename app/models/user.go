package models

import (
	"go-flashcard/app/database"
	"go-flashcard/app/utils"
	"log"
)

type User struct {
	ID       uint64
	Nickname string
	Email    string
	Password string
}

func (self *User) Validate(action string, form *utils.Form) bool {
	valid := true
	user := GetUserByEmail(form.Fields.Get("email"))

	switch action {
	case "create":
		if user != nil {
			valid = false
			form.Messages.AddMessage("Email already exists")
		}
		if form.Fields.Get("password") != form.Fields.Get("password2") {
			valid = false
			form.Messages.AddMessage("Passwords didn't match")
		}
		if valid {
			form.Messages = utils.MessagesSuccess
			form.Messages.AddMessage("Your account has been created. You can now login")

			self.Nickname = form.Fields.Get("nickname")
			self.Email = form.Fields.Get("email")
			self.Password = form.Fields.Get("password")

		}
	case "login":
		if user == nil {
			valid = false
			form.Messages.AddMessage("Email doesn't exist")
		} else if user.Password != form.Fields.Get("password") {
			valid = false
			form.Messages.AddMessage("Password is incorrect. Please try again")
		}
	}

	return valid
}

func GetUserByEmail(p_email string) *User {
	var (
		user     *User = nil
		id       uint64
		nickname string
		email    string
		password string
	)
	sqlStatement := "SELECT * FROM users WHERE email=$1"
	rows, err := database.DB.Query(sqlStatement, p_email)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for rows.Next() {
		err = rows.Scan(&id, &nickname, &email, &password)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		user = &User{
			ID:       id,
			Nickname: nickname,
			Email:    email,
			Password: password,
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return user
}

func (self *User) Save() {
	sqlStatement := `
	INSERT INTO users (nickname, email, password)
	VALUES ($1, $2, $3)`

	_, err := database.DB.Exec(sqlStatement, self.Nickname, self.Email, self.Password)
	if err != nil {
		log.Fatal(err)
		return
	}
}
