package student

// Pupil представляет собой студента и содержит информацию о его имени, возрасте, курсе и готовности к экзамену.
type Pupil struct {
	Name           string
	Age            int
	Course         string
	ReadyForExam   bool
	StudentHistory History
	Teacher        Teacher
}

// History содержит информацию о домашних заданиях и посещаемости студента.
type History struct {
	Homeworks  int // количество домашних заданий
	Attendance int // посещаемость
}

// Teacher представляет собой учителя и содержит информацию об его IQ и возможности списать.
type Teacher struct {
	IQ        int  // ум учителя. Если глупый, то студент увы не пройдет!
	Cheatering bool // шанс на руку помощи от учителя для прохождения экзамена
}
