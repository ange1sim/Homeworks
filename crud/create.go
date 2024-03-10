package crud

import (
	"fmt"
	"hw/models"
)

func Create(key string, user models.User) {
	models.Database[key] = user
	fmt.Printf("Данные %s, созданы\n", key)
}
