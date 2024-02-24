package pkg

import (
	"fmt"
	"hw/utils"
)

func DayThree() {
	fmt.Print("Выберите задание:\n1. Калькулятор\n")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	
	case 1:
		var num1, num2 float64

		fmt.Println("Введите 1-е число: ")
		fmt.Scan(&num1)
		fmt.Println("Введите 2-е число: ")
		fmt.Scan(&num2)

		result := utils.Difference(num1, num2)
		fmt.Println("Разница чисел: ", result)

		result = utils.Division(num1, num2)
		fmt.Println("Частное чисел: ", result)

		result = utils.Multiply(num1, num2)
		fmt.Println("Произведение чисел: ", result)

		result = utils.Sum(num1, num2)
		fmt.Println("Сумма чисел: ", result)
	
	default:
		fmt.Println("Такого задания нет!")
	}
}
