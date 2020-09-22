package model

import "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"

//PolarCoordinate Clase que contiene las coordenadas polares de un punto.
type PolarCoordinate struct {
	shareddomain.Coordinate
	Radius int
	Grades int
}

//CartesianCoordinate Clase que contiene las coordenadas cartesianas de un punto.
type CartesianCoordinate struct {
	shareddomain.Coordinate
	X float64
	Y float64
}

//AddGrades Método para sumarle grados a los grados actuales.
func (p PolarCoordinate) AddGrades(quantity int) {
	p.Grades += quantity
}

//NormalizeGrades Método para normalizar los grados. Ej.: Si hay 380°, este método los convierte a 20°.
//Deprecated: No se usa porque varían los datos por culpa de la constante math.Pi para el cálculo de los radianes, etc.
func (p PolarCoordinate) NormalizeGrades() {
	p.Grades %= 360
}

//GetRadius Obtiene el radio
func (p PolarCoordinate) GetRadius() int {
	return p.Radius
}

//GetGrades Obtiene los grados
func (p PolarCoordinate) GetGrades() int {
	return p.Grades
}

//GetX Obtiene el radio
func (p CartesianCoordinate) GetX() float64 {
	return p.X
}

//GetY Obtiene los grados
func (p CartesianCoordinate) GetY() float64 {
	return p.Y
}
