package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Admin struct {
	ID        uuid.UUID `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	UserName  string      `json:"user_name"`
	Password  string      `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (receiver *Admin) AddNew(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection admins.go/27 %e", err)
		return
	}

	defer conn.Release()

	receiver.ID = uuid.New()
	receiver.Password, err = HashPassword(receiver.Password)
	if err != nil {
		log.Printf("Can't hash password admins.go/36 %e\n", err)
		return
	}

	_, err = pool.Exec(context.Background(), `insert into "users"(id, first_name, last_name, username, password_hash) values($1, $2, $3, $4, $5);`,
		receiver.ID, receiver.FirstName, receiver.LastName, receiver.UserName, receiver.Password)
	if err != nil {
		log.Printf("Can't add to db admins.go/43 %e\n", err)
		return
	}

	return
}

func GetAdmin(username string, pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	//err = pool.QueryRow(context.Background(), `Select *from users where login = ($1)`, login).Scan(
	//	&User.Id,
	//	&User.Name,
	//	&User.Surname,
	//	&User.LastName,
	//	&User.Login,
	//	&User.Password,
	//	&User.Phone,
	//	&User.Role,
	//	&User.Status,
	//	&User.Position,
	//	&User.StatusLine)
	//if err != nil {
	//	log.Printf("Can't scan %e", err)
	//}
	return
}
