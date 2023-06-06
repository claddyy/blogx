// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Title     string
	Body      string
	Likes     int32
	Views     int32
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Bio       sql.NullString
	Password  string
	Isadmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserFollower struct {
	FollowerID  uuid.UUID
	FollowingID uuid.UUID
}

type UserLike struct {
	UserID uuid.UUID
	PostID uuid.UUID
}
