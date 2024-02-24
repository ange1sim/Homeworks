package exam

import (
	"fmt"
	"hw/student"
	"math/rand"
)

// TakeExam представляет собой функцию для проведения экзамена.
func TakeExam(client student.Pupil) student.Pupil {
	// Проводим проверки и симулируем экзамен.
	if client.StudentHistory.Homeworks < 10 || client.StudentHistory.Attendance < 25 {
		fmt.Println("Увы, не хватает домашних заданий или посещаемости. Экзамен провален!")
		client.ReadyForExam = false
		return client
	}

	if client.Teacher.IQ < 80 {
		fmt.Println("Увы, учитель слишком глупый. Экзамен провален!")
		client.ReadyForExam = false
		return client
	}

	// Симулируем шанс списать.
	if rand.Float32() < 0.2 && client.Teacher.Cheatering {
		fmt.Println("Учитель помог списать. Экзамен чудом пройден!")
		return client
	}

	fmt.Println("Экзамен успешно сдан!")
	return client
}
