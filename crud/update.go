package crud

import (
	"fmt"
	"hw/models"
)

func Update(key string, user models.User) {
	models.Database[key] = user
	fmt.Printf("Данные %s обновлены\n", key)
}
