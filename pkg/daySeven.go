package pkg

import "fmt"

func Reverse() {
	s := []interface{}{"Hello", "how", "are", "you", "?"}
	fmt.Println("Начальный массив: ", s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println("Измененный массив: ", s)
}

func EvenOdd(arr []int) ([]int, []int) {
	fmt.Println("Полученный массив: ", arr)
	evenSlice := make([]int, 0)
	oddSlice := make([]int, 0)

	for i, v := range arr {
		if i%2 == 0 {
			evenSlice = append(evenSlice, v)
		} else {
			oddSlice = append(oddSlice, v)
		}
	}
	return evenSlice, oddSlice
}
