// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"
)

type ExpiresIn string

const (
	ExpiresInNever ExpiresIn = "never"
	ExpiresIn1h    ExpiresIn = "1h"
	ExpiresIn2h    ExpiresIn = "2h"
	ExpiresIn10h   ExpiresIn = "10h"
	ExpiresIn1d    ExpiresIn = "1d"
	ExpiresIn2d    ExpiresIn = "2d"
	ExpiresIn1w    ExpiresIn = "1w"
	ExpiresIn1m    ExpiresIn = "1m"
	ExpiresIn1y    ExpiresIn = "1y"
)

func (e *ExpiresIn) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ExpiresIn(s)
	case string:
		*e = ExpiresIn(s)
	default:
		return fmt.Errorf("unsupported scan type for ExpiresIn: %T", src)
	}
	return nil
}

type Paste struct {
	ID        string
	CreatedAt time.Time
	ExpiresIn ExpiresIn
	DeletesAt sql.NullTime
	Title     string
}

type Pasty struct {
	ID       string
	PasteID  string
	Title    string
	Content  string
	Language string
}

type User struct {
	ID           string
	CreatedAt    time.Time
	Username     string
	AvatarUrl    string
	Contributor  bool
	Supporter    int32
	ProviderName string
	ProviderID   string
}
