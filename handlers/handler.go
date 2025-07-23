package handlers

import (
	"anime/config"
	"anime/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type AnimeView struct {
	Title    string
	Author   string
	CoverUrl string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "main.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cookie, err := r.Cookie("session_id")
	email := ""
	if err == nil {
		email = cookie.Value
	}

	var animes []models.Anime
	if err := config.DB.Find(&animes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Email  string
		Animes []models.Anime
	}{
		Email:  email,
		Animes: animes,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	tmplPath := filepath.Join("templates", "detail.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var animes []models.Anime
	if err := config.DB.First(&animes, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, animes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	email := cookie.Value
	tmplPath := filepath.Join("templates", "main.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, map[string]string{
		"Email": email,
	})
}
