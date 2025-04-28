package pkg

import (
	"github.com/google/uuid"
)

type UserAccount struct {
	ID         uuid.UUID `db:"id"`
	TelegramID int       `db:"telegram"`
	Balance    int       `db:"balance"`
	Blocked    bool      `db:"blocked"`
	Count      bool      `db:"count"`
	Social     string    `db:"social_network"`
	Url        string    `db:"url"`
}
