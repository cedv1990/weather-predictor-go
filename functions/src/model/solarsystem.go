package model

import (
	vo "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
)

//SolarSystem Entidad encargada del encapsulamiento de las propiedades de cada estrella.
type SolarSystem struct {
	Sun             *Star
	Days            []*Weather
	DaysWithMaxRain []int
	MaxPerimeter    float64
	DryDays         int
	RainyDays       int
	OptimalDays     int
	NormalDays      int
}

//NewSolarSystem Constructor de la clase SolarSystem
func NewSolarSystem(days int) *SolarSystem {
	solar := new(SolarSystem)

	solar.createPrediction(days)
	solar.calculateRelevanDataFromDays()

	return solar
}

//Agrega las predicciones a la lista de días que las contiene.
func (s SolarSystem) createPrediction(days int) {
	for i := 0; i < days; i++ {
		s.Days = append(s.Days, NewWeather(*s.Sun, i))
	}
}

//Método para extraer los datos relevantes de los días calculados.
func (s SolarSystem) calculateRelevanDataFromDays() {
	/**
	 * Se filtran los días por su condición climática.
	 */
	rainyDays := s.filterByWeatherCondition(vo.Rain)
	dryDays := s.filterByWeatherCondition(vo.Dry)
	optimalDays := s.filterByWeatherCondition(vo.Optimal)

	/**
	 * Se obtiene el perímetro máximo de todos los días de lluvia.
	 */
	s.MaxPerimeter = findMaxPerimeter(rainyDays)

	/**
	 * Se filtran los días de lluvia que tengan ese valor de perímetro máximo.
	 * Luego, se extrae solo el número de día.
	 */
	s.DaysWithMaxRain = getDaysWithMaxPerimeter(rainyDays, s.MaxPerimeter)

	/**
	 * Se asignan los totales.
	 */

	s.RainyDays = len(rainyDays)

	s.DryDays = len(dryDays)

	s.OptimalDays = len(optimalDays)

	s.NormalDays = len(s.Days) - (len(rainyDays) + len(dryDays) + len(optimalDays))
}

//Método para filtrar los días por su condición climática.
func (s SolarSystem) filterByWeatherCondition(condition vo.WeatherCondition) (ret []*Weather) {
	for _, o := range s.Days {
		if o.WeatherCondition == condition {
			ret = append(ret, o)
		}
	}
	return
}

//Método para encontrar el perímetro máximo de una lista de Weather
func findMaxPerimeter(a []*Weather) (max float64) {
	max = a[0].Perimeter
	for _, o := range a {
		if o.Perimeter > max {
			max = o.Perimeter
		}
	}
	return
}

//Método para obtener la lista de días que tienen el perímetro máximo
func getDaysWithMaxPerimeter(a []*Weather, p float64) (ret []int) {
	for _, o := range a {
		if o.Perimeter == p {
			ret = append(ret, o.Day)
		}
	}
	return
}
