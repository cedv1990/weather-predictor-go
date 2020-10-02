package model

import (
	vo "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	utils "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

//Weather Clase encargada del encapsulamiento de los datos correspondientes a un día específico en la predicción.
type Weather struct {
	Betasoide			Star
	Vulcano				Star
	Ferengi				Star
	Perimeter			float64
	WeatherCondition	vo.WeatherCondition
	Day					int

	sun 				Sun
}

//NewWeather Constructor de la clase Weather
func NewWeather(sun Sun, day int) *Weather {
	w := new(Weather)
	w.sun = sun
	w.Day = day
	w.WeatherCondition = vo.Normal
	w.initializeStars()
	w.setPositionByDayNumber()
	return w
}

//initializeStars Método para inicializar los planetas con su respectiva configuración.
func (w *Weather) initializeStars() {
	//Se instancian los planetas.
	w.Betasoide	= NewBetasoide()
	w.Vulcano	= NewVulcano()
	w.Ferengi	= NewFerengi()
}

//setPositionByDayNumber Método para asignar las posiciones correspondientes de los planetas dependiendo del número de día.
func (w *Weather) setPositionByDayNumber() {
	//Llamado al método que realiza el cálculo de la posición del planeta con respecto al número de día.
	w.Betasoide.SetPositionByDayNumber(w.Day)
	w.Vulcano.SetPositionByDayNumber(w.Day)
	w.Ferengi.SetPositionByDayNumber(w.Day)

	//Llamado al cálculo de la condición climática del día dependiendo de la posición de los planetas.
	w.setWeatherCondition()
}

//setWeatherCondition Método encargado de asignar la condición climática del día.
func (w *Weather) setWeatherCondition() {
	//Se obtienen las coordenadas cartesianas de cada planeta a partir de sus coordenadas polares.
	betasoideX, betasoideY	:= calculateCartesianCoordinateFromStar(w.Betasoide)
	vulcanoX, vulcanoY		:= calculateCartesianCoordinateFromStar(w.Vulcano)
	ferengiX, ferengiY		:= calculateCartesianCoordinateFromStar(w.Ferengi)
	sunX, sunY				:= calculateCartesianCoordinateFromStar(&w.sun)

	//Se crean las instancias
	betasoideCoordinate := CartesianCoordinate	{ X: betasoideX,Y: betasoideY }
	vulcanoCoordinate	:= CartesianCoordinate	{ X: vulcanoX, 	Y: vulcanoY }
	ferengiCoordinate	:= CartesianCoordinate	{ X: ferengiX, 	Y: ferengiY }
	sunCoordinate		:= CartesianCoordinate	{ X: sunX, 		Y: sunY }

	//Se realiza el llamado del método que devuelve la función que calcula la pendiente de la recta
	//formada a partir del planeta más lejano (Betasoide).
	functionToCalculateSlope := utils.GenerateFunctionToCalculateSlope(&betasoideCoordinate)

	//Se calculan las pendientes de las 3 rectas formadas a partir del planeta más lejano (Betasoide).
	slopeBetasoideFerengi	:= functionToCalculateSlope(&ferengiCoordinate)
	slopeBetasoideVulcano 	:= functionToCalculateSlope(&vulcanoCoordinate)
	slopeBetasoideSun 		:= functionToCalculateSlope(&sunCoordinate)

	//Se comparan las pendientes entre betasoide-ferengi y betasoide-vulcano.
	//Si son iguales, quiere decir que los 3 planetas están alineados.
	if slopeBetasoideFerengi == slopeBetasoideVulcano {
		/**
		 * Se compara la pendiente de cualquier recta formada entre los planetas con la pendiente
		 * formada entre el planeta más lejano (Betasoide) y el Sol.
		 * Si son iguales, quiere decir que los 3 planetas están alineados con el sol, lo cual
		 * significa que habrá sequía.
		 * Si son diferentes, quiere decir que los 3 planetas están alineados pero no con el sol, lo
		 * cual significa que habrán condiciones óptimas de presión y temperatura.
		 */
		if slopeBetasoideFerengi == slopeBetasoideSun {
			w.WeatherCondition = vo.Dry
		} else {
			w.WeatherCondition = vo.Optimal
		}
	} else {
		/**
		 * Al no estár alineados, los planetas forman un triángulo.
		 * Para calcular el perímetro de cualquier triángulo, se requiere conocer la longitud de sus lados,
		 * para así sumarlos y hallar el valor.
		 */

		//Se calculan las distancias entre los 3 planetas.
		distanceBetasoideFerengi	:= utils.GetDistanceBetweenPoints(&betasoideCoordinate, &ferengiCoordinate)
		distanceBetasoideVulcano 	:= utils.GetDistanceBetweenPoints(&betasoideCoordinate, &vulcanoCoordinate)
		distanceVulcanoFerengi 		:= utils.GetDistanceBetweenPoints(&vulcanoCoordinate, &ferengiCoordinate)

		/**
		 * Se calcula el perímetro del triángulo, mediante la suma de sus lados.
		 * Posteriormente, se revisará en la clase {@link SolarSystem} el perímetro más alto,
		 * lo cual servirá para detectar los días con picos altos de lluvia; éstos corresponden
		 * a aquellos días que compartan ese valor del perímetro mayor.
		 */
		w.Perimeter = distanceBetasoideFerengi + distanceBetasoideVulcano + distanceVulcanoFerengi

		/**
		 * Se realiza el llamado al método encargado de revisar si el sol se encuentra dentro del triángulo
		 * formado entre los 3 planetas.
		 * Si lo está, quiere decir que la condición del clima será de lluvia.
		 */
		if utils.EvaluateIfPointIsInsideTheTriangle(&betasoideCoordinate, &ferengiCoordinate, &vulcanoCoordinate, &sunCoordinate) {
			w.WeatherCondition = vo.Rain
		}
	}
}

//calculateCartesianCoordinateFromStar Método para calcular la coordenada cartesiana a partir de la coordenada polar.
func calculateCartesianCoordinateFromStar(star Star) (x, y float64) {
	return utils.GetCartesianCoordinateFromPolarCoordinate(star.GetPolarCoordinate())
}
