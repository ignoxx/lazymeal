module lazymeal

go 1.23

toolchain go1.23.0

// uncomment for local development on the superkit core.
// replace github.com/anthdm/superkit => ../

require (
	github.com/a-h/templ v0.2.778
	github.com/anthdm/superkit v0.0.0-20240701091803-e7f8e0aad3e9
	github.com/go-chi/chi/v5 v5.1.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.23
	golang.org/x/crypto v0.27.0
	golang.org/x/image v0.20.0
	gorm.io/gorm v1.25.11
)

require (
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.4.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.18.0 // indirect
)
