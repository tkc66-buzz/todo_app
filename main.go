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
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

}
