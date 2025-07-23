package handlers

import (
	"html/template"
	"net/http"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email обязателен", http.StatusBadRequest)
			return
		}

		// Устанавливаем cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    email, // в реальности лучше использовать UUID/token
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
