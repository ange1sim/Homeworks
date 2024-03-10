package crud

import (
	"fmt"
	"hw/models"
)
func Read(key string) {
	user, ok := models.Database[key]
	if ok {
		fmt.Printf("Username: %s, Email: %s\n", user.Username, user.Email)
	} else {
		fmt.Printf("Пользователь с именем %s не найден\n", key)
	}
}