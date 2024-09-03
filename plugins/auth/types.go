package auth

import (
	"context"
	"lazymeal/app/db"
	"lazymeal/app/db/sqlc"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Event name constants
const (
	UserSignupEvent         = "auth.signup"
	ResendVerificationEvent = "auth.resend.verification"
)

// UserWithVerificationToken is a struct that will be sent over the
// auth.signup event. It holds the User struct and the Verification token string.
type UserWithVerificationToken struct {
	User  sqlc.User
	Token string
}

type Auth struct {
	UserID   uint
	Email    string
	LoggedIn bool
}

func (auth Auth) Check() bool {
	return auth.LoggedIn
}

func createUserFromFormValues(values SignupFormValues) (sqlc.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(values.Password), bcrypt.DefaultCost)
	if err != nil {
		return sqlc.User{}, err
	}

	user, err := db.Get().CreateUser(context.TODO(), sqlc.CreateUserParams{
		Email:        values.Email,
		FirstName:    values.FirstName,
		LastName:     values.LastName,
		PasswordHash: string(hash),
	})

	return user, err
}

type Session struct {
	UserID    uint
	Token     string
	IPAddress string
	UserAgent string
	ExpiresAt time.Time
	CreatedAt time.Time
	User      sqlc.User
}
