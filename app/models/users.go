package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() error {
	cmd := `insert into users (
	uuid,
	name,
	email,
	password,
	created_at) values (?, ?, ?, ?, ?)`

	_, err := Db.Exec(cmd, createUUID(), u.Name, u.Email,
		Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (User, error) {
	user := User{}
	cmd := `
	select id, uuid, name, email, password, created_at
	from users where id = ?
	`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() error {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
