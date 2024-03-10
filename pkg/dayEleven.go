package pkg

import (
	"fmt"
	"strings"

	"github.com/vodolaz095/caesar"
)

func stringToBytes(s string) []uint8 {
	return []uint8(s)
}

func bytesToString(b []uint8) string {
	return string(b)
}

func DayEleven() {
	var choice int
	fmt.Println("Выберите задачу:")
	fmt.Println("1. Переобразование строк и слайса")
	fmt.Println("2. Шифр Цезаря")
	fmt.Println("3. Редактирование строки")
	fmt.Print("Введите номер задачи: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var inputString string
		fmt.Print("Введите строку: ")
		fmt.Scan(&inputString)
		byteSlice := stringToBytes(inputString)
		fmt.Println(byteSlice)
		fmt.Println(bytesToString(byteSlice))
	case 2:
		cryptor := caesar.NewRussian()
		fmt.Println("Введите слово:")
		var word string
		fmt.Scan(&word)
		fmt.Println(cryptor.Decrypt(word, 3))
	case 3:
		text := "Мы+получили+текст,+в+котором+по+какой-то+ошибке+все+пробелы+заменились+на+плюсы.+Надо+это+исправить"
		fmt.Println(strings.Split(text, "+"))
	default:
		fmt.Println("Неверный выбор задачи. Попробуйте снова.")
	}
}
