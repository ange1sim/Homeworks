package main

import (
	"fmt"
	"hw/pkg"
	"log"
)

func main() {

	log.Println("__________________________________Запуск программы__________________________________")
	fmt.Println("Здравствуйте! Пожалуйста, выберите день обучения:	")
	fmt.Println("1. День 1-й")
	fmt.Println("2. День 2-й")
	fmt.Println("3. День 3-й")
	fmt.Println("4. День 4-й")
	fmt.Println("5. День 5-й")
	fmt.Println("6. День 6-й")
	fmt.Println("7. День 7-й")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		pkg.DayOne()
	case 2:
		pkg.DayTwo()
	case 3:
		pkg.DayThree()
	case 4:
		pkg.DayFour()
	case 5:
		pkg.DayFive()
	case 6:
		pkg.DaySix()
	case 7:
		fmt.Println("Выберите задание:")
		fmt.Println("1. Reverse array")
		fmt.Println("2. Разбиение массива на слайсы")
		var choice int
		fmt.Scan(&choice)
		if choice == 1{
			pkg.Reverse()
		} else if choice == 2{
			array := []int{1, 2, 3, 4, 5}
    		even, odd := pkg.EvenOdd(array)
    		fmt.Printf("Четные числа: %d\nНечётные числа: %d\n", odd, even)
		}
	default:
		fmt.Println("Вы выбрали неверный день обучения")
	}
	log.Println("_______________________________Программа завершена :)_______________________________")
}
