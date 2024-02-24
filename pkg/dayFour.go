package pkg

import (
	"fmt"
	"hw/utils"
)

func DayFour() {
	fmt.Println("Выберите задание:")
	fmt.Println("1. Единая программа")
	fmt.Println("2. Бoнусное задание: таблица умножения")
	fmt.Println("3. Бoнусное задание: Фибоначчи")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Вы уже наблюдаете это дз!")
	case 2:
		for i := 2; i <= 9; i++ {
			for x := 2; x <= 9; x++ {
				fmt.Printf("%v * %v = %v   \t", x, i, i*x)
			}
			fmt.Printf("\n")
		}
	case 3:
		var num uint
		fmt.Println("Введите число: ")
		fmt.Scan(&num)

		result := utils.Fibonacci(num)
		fmt.Println(result)
	default:
		fmt.Println("Неверно")
		return
	}
}
