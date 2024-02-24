package pkg

import (
	"fmt"
	"hw/exam"
	"hw/student"
	"log"
	"math/rand"
)

// DayFive представляет собой функцию для симуляции экзамена.
func DayFive() {
	// Создаем студента с случайными параметрами.
	var Student student.Pupil = student.Pupil{
		Name: "Сим Ангелина Дмитриевна", Age: rand.Intn(40), Course: "Golang-разработка",
		ReadyForExam: true, StudentHistory: student.History{Homeworks: rand.Intn(20) + 5, Attendance: rand.Intn(40) + 10},
		Teacher: student.Teacher{IQ: rand.Intn(50) + 50, Cheatering: rand.Float32() < 0.3}, // случайные параметры учителя
	}

	// Проверка возраста студента.
	if Student.Age < 18 {
		fmt.Printf("Увы, я пока еще малышарик :( %+v", Student)
	} else if Student.Age > 30 {
		fmt.Printf("Мне пора на пенсию %+v", Student)
	}

	log.Println("Я готов к экзамену!")
	log.Println("Я зашел в офис!")

	// Вызываем функцию для проведения экзамена.
	Student = exam.TakeExam(Student)

	// Проверка результата после экзамена.
	if !Student.ReadyForExam {
		log.Printf("Увы, я не был хорошо подготовлен к экзамену: %+v", Student)
	} else if Student.Teacher.IQ < 80 {
		fmt.Println("Увы, учитель не смог хорошо подготовить нас :(")
	} else if Student.Teacher.Cheatering {
		fmt.Println("Учителю Респект!!!!!!!!!")
	} else {
		log.Printf("Я прошел экзамен! И чувствую себя прекрасно!: %+v", Student)
	}
}
