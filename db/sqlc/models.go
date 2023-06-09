// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID
	UserID            uuid.NullUUID
	Type              sql.NullString
	Provider          sql.NullString
	ProviderAccountID sql.NullString
	RefreshToken      sql.NullString
	AccessToken       sql.NullString
	ExpiresAt         sql.NullTime
	TokenType         sql.NullString
	Scope             sql.NullString
	IDToken           sql.NullString
	SessionState      sql.NullString
}

type User struct {
	ID            uuid.UUID
	Username      sql.NullString
	Email         sql.NullString
	EmailVerified sql.NullTime
	Image         sql.NullString
}

type Userrefreshtoken struct {
	ID           uuid.UUID
	UserID       uuid.NullUUID
	RefreshToken sql.NullString
}
