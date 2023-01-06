package models

import (
	"log"
	"time"
)

type ToDo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateToDo(content string) error {
	cmd := `
	insert into todos
	(content, user_id, created_at) values (?, ?, ?)
	`
	_, err := Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetToDo(id int) (ToDo, error) {
	todo := ToDo{}
	cmd := `
	select id, content, user_id, created_at
	from todos where id = ?
	`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}
