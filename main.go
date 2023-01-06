package main

import (
	"fmt"

	"github.com/tkc66-buzz/todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{
		Name:     "test",
		Email:    "test@example.com",
		Password: "testPassword123!",
	}
	fmt.Println(u)
	u.CreateUser()
}
