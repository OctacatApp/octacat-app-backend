// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package domain

import (
	"time"
)

type Chat struct {
	ID         string
	FromUserID string
	ToUserID   string
	MessageID  string
	CreatedAt  time.Time
	CreatedBy  string
	UpdatedAt  time.Time
	UpdatedBy  string
	DeletedAt  time.Time
	DeletedBy  string
	IsDeleted  bool
}

type ChatMessage struct {
	ID        string
	Message   string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
	IsDeleted bool
}

type User struct {
	ID           string
	Name         string
	Email        string
	Password     string
	ProfileImage string
	CreatedAt    time.Time
	CreatedBy    string
	UpdatedAt    time.Time
	UpdatedBy    string
	DeletedAt    time.Time
	DeletedBy    string
	IsDeleted    bool
}
