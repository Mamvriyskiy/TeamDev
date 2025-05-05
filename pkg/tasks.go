package pkg

type NewTasks struct {
	TelegramID int       `db:"telegram"`
	Balance    int       `db:"balance"`
	Blocked    bool      `db:"blocked"`
	Count      bool      `db:"count"`
	Social     string    `db:"social_network"`
	Url        string    `db:"url"`
}
