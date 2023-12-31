// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      *string   `json:"name"`
	Email     *string   `json:"email"`
	LastName  *string   `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
}
