package dto

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt null.Time `json:"updatedAt"`
	DeletedAt null.Time `json:"deletedAt"`
}
