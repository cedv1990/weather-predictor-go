package valueobjects

import value "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"

//Weather Clase encargada del encapsulamiento de los datos correspondientes a un día específico en la predicción.
type Weather struct {
	Betasoide value.Star
	Vulcano   value.Star
	Ferengi   value.Star

	Perimeter        int
	WeatherCondition value.WeatherCondition
}
