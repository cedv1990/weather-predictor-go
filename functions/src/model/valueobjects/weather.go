package valueobjects

import (
	model "github.com/cedv1990/weather-predictor-go/functions/src/model"
	value "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	utils "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

//Weather Clase encargada del encapsulamiento de los datos correspondientes a un día específico en la predicción.
type Weather struct {
	Betasoide        *value.Star
	Vulcano          *value.Star
	Ferengi          *value.Star
	Perimeter        int
	WeatherCondition value.WeatherCondition
	Day              int

	sun *value.Star
}

//NewWeather Constructor de la clase Weather
func NewWeather(sun value.Star, day int) *Weather {
	w := new(Weather)
	w.initializeStars()
	w.setPositionByDayNumber()
	return w
}

func (w Weather) initializeStars() {
	w.Betasoide = value.NewStar("Betasoide", 2000, 0, 3, true)
	w.Vulcano = value.NewStar("Vulcano", 1000, 0, 5, false)
	w.Ferengi = value.NewStar("Ferengi", 500, 0, 1, true)
}

func (w Weather) setPositionByDayNumber() {
	w.Betasoide.SetPositionByDayNumber(w.Day)
	w.Vulcano.SetPositionByDayNumber(w.Day)
	w.Ferengi.SetPositionByDayNumber(w.Day)

	w.setWeatherCondition()
}

func (w Weather) setWeatherCondition() {
	betasoideCartesianCoordinate := w.calculateCartesianCoordinateFromStar(w.Betasoide)
	vulcanoCartesianCoordinate := w.calculateCartesianCoordinateFromStar(w.Vulcano)
	ferengiCartesianCoordinate := w.calculateCartesianCoordinateFromStar(w.Ferengi)
	sunCartesianCoordinate := w.calculateCartesianCoordinateFromStar(w.sun)
}

func (w Weather) calculateCartesianCoordinateFromStar(star *value.Star) model.CartesianCoordinate {
	/**
	 * Llamado del método de cálculo presente en utils.go.
	 */
	return utils.GetCartesianCoordinateFromPolarCoordinate(star.PolarCoordinate)
}
