package pkg

import (
	"fmt"
	"hw/operations"
)

func DaySeven() {
	fmt.Println("Выберите задание:")
	fmt.Println("1. Reverse array")
	fmt.Println("2. Разбиение массива на слайсы")
	var choice int
	fmt.Scan(&choice)
	if choice == 1 {
		operations.Reverse()
	} else if choice == 2 {
		array := []int{1, 2, 3, 4, 5}
		even, odd := operations.EvenOdd(array)
		fmt.Printf("Четные числа: %d\nНечётные числа: %d\n", odd, even)
	}
}
