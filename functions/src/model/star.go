package model

type (
	//Star Interface para la herencia y envío de info a otros métodos.
	Star interface {
		SetPositionByDayNumber(dayNumber int)
		GetPolarCoordinate() PolarCoordinate
	}

	//BaseStar Entidad encargada del encapsulamiento de las propiedades de cada estrella.
	BaseStar struct {
		Star
		PolarCoordinate *PolarCoordinate
		Name            string
		Grades          int
		Velocity        int
		Clockwise       bool
	}

	Betasoide struct {
		BaseStar
	}

	Vulcano struct {
		BaseStar
	}

	Ferengi struct {
		BaseStar
	}

	Sun struct {
		BaseStar
	}
)

//fillData Constructor de la clase BaseStar
func (star *BaseStar) fillData(name string, distance, grades, velocity int, clockwise bool) {
	star.Name = name
	star.PolarCoordinate = &PolarCoordinate{Grades: grades, Radius: distance}
	star.Velocity = velocity
	star.Clockwise = clockwise
}

//SetPositionByDayNumber Método encargado de asignar la nueva coordenada polar de la estrella a partir de un número de día.
func (star *BaseStar) SetPositionByDayNumber(dayNumber int) {
	if star.Clockwise {
		star.PolarCoordinate.AddGrades(-star.Velocity * dayNumber)
	} else {
		star.PolarCoordinate.AddGrades(star.Velocity * dayNumber)
	}
}

//GetPolarCoordinate Obtiene la coordenada polar
func (star *BaseStar) GetPolarCoordinate() PolarCoordinate {
	return *star.PolarCoordinate
}

//NewBetasoide constructor de Betasoide.
func NewBetasoide() *Betasoide {
	beta := &Betasoide{}
	beta.fillData("Betasoide", 2000, 0, 3, true)
	return beta
}

//NewVulcano constructor de Vulcano.
func NewVulcano() *Vulcano {
	vul := &Vulcano{}
	vul.fillData("Vulcano", 1000, 0, 5, false)
	return vul
}

//NewFerengi constructor de Ferengi.
func NewFerengi() *Ferengi {
	fer := &Ferengi{}
	fer.fillData("Ferengi", 500, 0, 1, true)
	return fer
}

//NewSun constructor de Sun.
func NewSun() *Sun {
	sun := &Sun{}
	sun.fillData("Sun", 0, 0, 0, false)
	return sun
}