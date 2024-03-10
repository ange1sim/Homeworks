package pkg

import (
	"fmt"
	"hw/crud"
	"hw/models"
)

func DayEight() {
	var name, email string
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите email: ")
	fmt.Scan(&email)
	user1 := models.User{Username: name, Email: email}
	crud.Create(name, user1)
	crud.Read(name)

	user1.Email = email
	crud.Update(name, user1)

	crud.Delete(name)
	crud.Read(name)
}
