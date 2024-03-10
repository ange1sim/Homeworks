package crud

import (
	"fmt"
	"hw/models"
)
func Delete(key string) {
	delete(models.Database, key)
	fmt.Printf("Данные %s удалены\n", key)
}
