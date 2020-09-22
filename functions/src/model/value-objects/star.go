package valueobjects

//Star Entidad encargada del encapsulamiento de las propiedades de cada estrella.
type Star struct {
	PolarCoordinate PolarCoordinate
	Name            string
	Grades          int
	Velocity        int
	Clockwise       bool
}

//NewStar Constructor de la clase Star
func NewStar(name string, distance, grades, velocity int, clockwise bool) *Star {
	star := new(Satr)
	star.Name = name
	star.PolarCoordinate = PolarCoordinate{Grades: grades, Radius: distance}
	star.Velocity = velocity
	star.Clockwise = clockwise

	return star
}

//SetPositionByDayNumber Método encargado de asignar la nueva coordenada polar de la estrella a partir de un número de día.
func (star NewStar) SetPositionByDayNumber(dayNumber int) {
	if star.Clockwise {
		star.PolarCoordinate.AddGrades(-star.Velocity * dayNumber)
	} else {
		star.PolarCoordinate.AddGrades(star.Velocity * dayNumber)
	}
}
