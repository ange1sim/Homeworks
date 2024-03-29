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
	fmt.Println("8,9. День 8, 9-й")
	fmt.Println("10. День 10-й")
	fmt.Println("11. День 11-й")
	fmt.Println("12. День 12-й")
	fmt.Println("13. День 13-й")

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
		pkg.DaySeven()
	case 8, 9:
		pkg.DayEight()
	case 10,12:
		pkg.DayTen()
	case 11:
		pkg.DayEleven()
	case 13:
		pkg.DayThirteenth()
	default:
		fmt.Println("Вы выбрали неверный день обучения")
	}
	log.Println("_______________________________Программа завершена :)_______________________________")
}
