package controllers

import (
	"go-flashcard/app/models"
	"go-flashcard/app/templates"
	"go-flashcard/app/utils"
	"log"
	"net/http"
)

func UserHandler(fn func(http.ResponseWriter, *http.Request, *utils.Form)) http.HandlerFunc {
	form := &utils.Form{}
	form.Messages = &utils.Messages{}

	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
			return
		}
		form.Fields = r.Form

		fn(w, r, form)
	}
}

func ShowRegister(w http.ResponseWriter, r *http.Request, form *utils.Form) {
	templates.RenderTemplate(w, "register", form)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, form *utils.Form) {
	user := &models.User{}
	form.Messages.ClearMessage()
	form.Messages = utils.MessagesDanger
	valid := user.Validate("create", form)

	if !valid {
		templates.RenderTemplate(w, "register", form)
	} else {
		user.Save()
		http.Redirect(w, r, "/users/login", http.StatusFound)
	}
}

func ShowLogin(w http.ResponseWriter, r *http.Request, form *utils.Form) {
	form.Messages = utils.MessagesSuccess
	templates.RenderTemplate(w, "login", form)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, form *utils.Form) {
	user := &models.User{}
	form.Messages.ClearMessage()
	form.Messages = utils.MessagesDanger
	valid := user.Validate("login", form)

	if !valid {
		templates.RenderTemplate(w, "login", form)
	} else {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	}
}
