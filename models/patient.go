package models

type Patient struct {
	Name    string
	Age     int
	Order   int
	Organs  Organs
	MedBook MedBook
}
