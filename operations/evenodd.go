package operations

import "fmt"

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
