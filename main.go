package main

import (
	"fmt"

	"github.com/tkc66-buzz/todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{
	// 	Name:     "test",
	// 	Email:    "test@example.com",
	// 	Password: "testPassword123!",
	// }
	// fmt.Println(u)
	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateToDo("First ToDo")
	todo, _ := models.GetToDo(1)
	fmt.Println(todo)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

}
