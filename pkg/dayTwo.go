package pkg

import (
	"fmt"
	"math"
)

func DayTwo() {
	fmt.Println("Выберите задание:")
	fmt.Println("1. Периметр квадрата")
	fmt.Println("2. Периметр прямоугольника")
	fmt.Println("3. Площадь прямоугольника")
	fmt.Println("4. Бoнусное задание: Юбилей")

	var choice, num, num2 int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Ведите сторону квадрата:")
		fmt.Scan(&num)

		fmt.Println("Периметр квадрата равен: ", num*4)
	case 2:

		fmt.Println("Введите ширину:")
		fmt.Scan(&num)
		fmt.Println("Введите длину:")
		fmt.Scan(&num2)

		fmt.Println("периметр прямоугольника =", (num+num2)*2)
	case 3:

		fmt.Println("Введите ширину:")
		fmt.Scan(&num)
		fmt.Println("Введите длину:")
		fmt.Scan(&num2)

		fmt.Println("периметр прямоугольника =", num*num2)
	case 4:
		var num, num2 float64
		fmt.Println("Ведите ваш возраст:")
		fmt.Scan(&num)
		num = (num / 10)
		num2 = math.Ceil(num)
		num2 = (num2 * 10) - (num * 10)
		fmt.Println("До юбилея вам осталось:", num2)
	default:
		fmt.Println("Данного задания нет!")
	}
}