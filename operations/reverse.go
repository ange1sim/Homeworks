package operations

import "fmt"

func Reverse() {
	s := []interface{}{"Hello", "how", "are", "you", "?"}
	fmt.Println("Начальный массив: ", s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println("Измененный массив: ", s)
}
