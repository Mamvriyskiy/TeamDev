package repository

import (
	"fmt"
	"log"

	"github.com/Mamvriyskiy/TeamDev/pkg"
	"github.com/google/uuid" // Для генерации UUID
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) AddSocialUser(userID int, url string) error {
	var us_id string
	query := `select id from user_account where telegram = $1`
	err := r.db.Get(&us_id, query, userID)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(us_id)
	query = `INSERT INTO account (id, owner_id, social_network, url) VALUES ($1, $2, $3, $4)`
	_ = r.db.QueryRow(query, uuid.New(), us_id, "vk", url)

	return nil
}

func (r *UserPostgres) RegisterUser(userID int) (result string, err error) {
	user := pkg.UserAccount{
		ID:         uuid.New(), // Генерация нового UUID
		TelegramID: userID,
		Balance:    100,
		Blocked:    false,
	}

	// Вставка данных в таблицу user_account
	_, err = r.db.NamedExec(`INSERT INTO user_account (id, telegram, balance, blocked) VALUES (:id, :telegram, :balance, :blocked)`, &user)
	if err != nil {
		log.Fatalln(err)
	}

	return "", nil
}

func (r *UserPostgres) ProfileUser(userID int) (user pkg.UserAccount, err error) {
	query := `SELECT us.telegram, us.balance, COUNT(o.status), ac.url, ac.social_network 
		FROM user_account us
		JOIN offer o ON us.id = o.id
		join account ac on ac.owner_id = us.id
		WHERE us.telegram = $1 AND o.status = 'active'
		GROUP BY us.telegram, us.balance, ac.url, ac.social_network;`

	var s []pkg.UserAccount
	err = r.db.Select(&s, query, userID)
	fmt.Println(err, s)
	if err != nil {
		log.Println("Error", "Select", "Error get user profile:", err, userID)
		return pkg.UserAccount{}, err
	}

	return user, nil
}

func (r *UserPostgres) CreateTasks(userID int, tasks pkg.NewTasks) (err error) {
	return nil
}

func (r *UserPostgres) CheckStatusTasks(userID int, nameTasks string) (count, all int, err error) {
	return 0, 0, nil
}

