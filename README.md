# Anime Catalog Web App

This is a simple Go-based web application for browsing anime. Users can log in with an email (mock login), view a list of anime titles, and see detailed information for each title.

## Features

- 🧾 Anime catalog with cover image, title, and author
- 🔍 Detail page for each anime
- 📧 Fake login with email and session cookies
- 🗃️ PostgreSQL database (AWS RDS supported)
- 🌐 HTML templates rendered with Go's `html/template`
- 📁 Project structured with `handlers/`, `models/`, `middleware/`, and `routes/` folders

## Technologies Used

- Go (Golang)
- GORM ORM
- PostgreSQL
- Gorilla Mux
- HTML + CSS (Tailwind or custom styles)
- AWS RDS (for PostgreSQL hosting)

## Setup Instructions

1. Clone the repo:
   ```bash
   git clone https://github.com/MukhammedK/anime.git
   cd anime
