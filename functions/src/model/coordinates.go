package model

//Coordinate interfaz para que las coordenadas la implementen
type Coordinate interface {
	GetGrades() int
	GetRadius() int
	GetX() float64
	GetY() float64
}
