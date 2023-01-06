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

func (u *User) GetToDosByUser() ([]ToDo, error) {
	cmd := `
	select id, content, user_id, created_at from todos
	where user_id = ?
	`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	var todos []ToDo
	for rows.Next() {
		var todo ToDo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *ToDo) UpdateToDo() error {
	cmd := `
			update todos set content = ?, user_id = ?
			where id = ?
			`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *ToDo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
