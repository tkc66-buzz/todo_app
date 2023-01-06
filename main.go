package main

import (
	"fmt"

	"github.com/tkc66-buzz/todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{
	// 	Name:     "test2",
	// 	Email:    "test2@example.com",
	// 	Password: "testPassword123!",
	// }
	// fmt.Println(u)
	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateToDo("First ToDo")

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user2, _ := models.GetUser(3)
	// todos, _ := user2.GetToDosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetToDo(2)
	t.DeleteTodo()
	fmt.Println(t)
	newt, _ := models.GetToDo(3)
	fmt.Println(newt)
}
