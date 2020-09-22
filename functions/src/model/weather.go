package model

import (
	vo "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	utils "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

//Weather Clase encargada del encapsulamiento de los datos correspondientes a un día específico en la predicción.
type Weather struct {
	Betasoide        *Star
	Vulcano          *Star
	Ferengi          *Star
	Perimeter        float64
	WeatherCondition vo.WeatherCondition
	Day              int

	sun *Star
}

//NewWeather Constructor de la clase Weather
func NewWeather(sun Star, day int) *Weather {
	w := new(Weather)
	w.initializeStars()
	w.setPositionByDayNumber()
	return w
}

//Método para inicializar los planetas con su respectiva configuración.
func (w Weather) initializeStars() {
	/**
	 * Se instancian los planetas con:
	 * Nombre
	 * Distancia al sol
	 * Grados
	 * Velocidad (grados por día)
	 * Movimiento con respecto a las manecillas del reloj
	 */
	w.Betasoide = NewStar("Betasoide", 2000, 0, 3, true)
	w.Vulcano = NewStar("Vulcano", 1000, 0, 5, false)
	w.Ferengi = NewStar("Ferengi", 500, 0, 1, true)
}

//Método para asignar las posiciones correspondientes de los planetas dependiendo del número de día.
func (w Weather) setPositionByDayNumber() {
	/**
	 * Llamado al método que realiza el cálculo de la posición del planeta con respecto al número de día.
	 */
	w.Betasoide.SetPositionByDayNumber(w.Day)
	w.Vulcano.SetPositionByDayNumber(w.Day)
	w.Ferengi.SetPositionByDayNumber(w.Day)

	/**
	 * Llamado al cálculo de la condición climática del día dependiendo de la posición de los planetas.
	 */
	w.setWeatherCondition()
}

//Método encargado de asignar la condición climática del día.
func (w Weather) setWeatherCondition() {
	/**
	 * Se obtienen las coordenadas cartesianas de cada planeta a partir de sus coordenadas polares.
	 */
	betasoideCartesianCoordinateX, betasoideCartesianCoordinateY := w.calculateCartesianCoordinateFromStar(w.Betasoide)
	vulcanoCartesianCoordinateX, vulcanoCartesianCoordinateY := w.calculateCartesianCoordinateFromStar(w.Vulcano)
	ferengiCartesianCoordinateX, ferengiCartesianCoordinateY := w.calculateCartesianCoordinateFromStar(w.Ferengi)
	sunCartesianCoordinateX, sunCartesianCoordinateY := w.calculateCartesianCoordinateFromStar(w.sun)

	/**
	 *Se crean las instancias
	 */
	betasoideCartesianCoordinate := CartesianCoordinate{X: betasoideCartesianCoordinateX, Y: betasoideCartesianCoordinateY}
	vulcanoCartesianCoordinate := CartesianCoordinate{X: vulcanoCartesianCoordinateX, Y: vulcanoCartesianCoordinateY}
	ferengiCartesianCoordinate := CartesianCoordinate{X: ferengiCartesianCoordinateX, Y: ferengiCartesianCoordinateY}
	sunCartesianCoordinate := CartesianCoordinate{X: sunCartesianCoordinateX, Y: sunCartesianCoordinateY}

	/**
	 * Se realiza el llamado del método que devuelve la función que calcula la pendiente de la recta
	 * formada a partir del planeta más lejano (Betasoide).
	 */
	functionToCalculateSlope := utils.GenerateFunctionToCalculateSlope(&betasoideCartesianCoordinate)

	/**
	 * Se calculan las pendientes de las 3 rectas formadas a partir del planeta más lejano (Betasoide).
	 */
	slopeBetasoideFerengi := functionToCalculateSlope(&ferengiCartesianCoordinate)
	slopeBetasoideVulcano := functionToCalculateSlope(&vulcanoCartesianCoordinate)
	slopeBetasoideSun := functionToCalculateSlope(&sunCartesianCoordinate)

	/**
	 * Se comparan las pendientes entre betasoide-ferengi y betasoide-vulcano.
	 * Si son iguales, quiere decir que los 3 planetas están alineados.
	 */
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

		/**
		 * Se calculan las distancias entre los 3 planetas.
		 */
		distanceBetasoideFerengi := utils.GetDistanceBetweenPoints(&betasoideCartesianCoordinate, &ferengiCartesianCoordinate)
		distanceBetasoideVulcano := utils.GetDistanceBetweenPoints(&betasoideCartesianCoordinate, &vulcanoCartesianCoordinate)
		distanceVulcanoFerengi := utils.GetDistanceBetweenPoints(&vulcanoCartesianCoordinate, &ferengiCartesianCoordinate)

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
		if utils.EvaluateIfPointIsInsideTheTriangle(&betasoideCartesianCoordinate, &ferengiCartesianCoordinate, &vulcanoCartesianCoordinate, &sunCartesianCoordinate) {
			w.WeatherCondition = vo.Rain
		}
	}
}

//Método para calcular la coordenada cartesiana a partir de la coordenada polar.
func (w Weather) calculateCartesianCoordinateFromStar(star *Star) (x, y float64) {
	/**
	 * Llamado del método de cálculo presente en utils.go.
	 */
	return utils.GetCartesianCoordinateFromPolarCoordinate(star.PolarCoordinate)
}
