package model

//PolarCoordinate Clase que contiene las coordenadas polares de un punto.
type PolarCoordinate struct {
	Radius int
	Grades int
}

//CartesianCoordinate Clase que contiene las coordenadas cartesianas de un punto.
type CartesianCoordinate struct {
	X int
	Y int
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
