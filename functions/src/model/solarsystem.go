package model

import vo "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"

//SolarSystem Entidad encargada del encapsulamiento de los datos correspondientes al sistema solar
//que contiene la lista de todos los días de la predicción.
type SolarSystem struct {
	Sun             Sun			`json:"-"`
	Days            []*Weather	`json:"-"` //Lista de todos los días que se predijeron.
	DaysWithMaxRain []int		`json:"daysWithMaxRain"`
	MaxPerimeter    float64		`json:"maxPerimeter"`
	DryDays         int			`json:"dryDays"`
	RainyDays       int			`json:"rainyDays"`
	OptimalDays     int			`json:"optimalDays"`
	NormalDays      int			`json:"normalDays"`
}

//NewSolarSystem Constructor de la clase SolarSystem
func NewSolarSystem(days int) *SolarSystem {
	solar := new(SolarSystem)

	//Instancias
	solar.Days = []*Weather{}
	solar.Sun = *NewSun()
	solar.DaysWithMaxRain = []int{}
	solar.MaxPerimeter = -1
	solar.DryDays = 0
	solar.RainyDays = 0
	solar.OptimalDays = 0
	solar.NormalDays = 0

	//Ejecuciones
	solar.createPrediction(days)
	solar.calculateRelevanDataFromDays()

	return solar
}

//createPrediction Agrega las predicciones a la lista de días que las contiene.
func (s *SolarSystem) createPrediction(days int) {
	for i := 0; i < days; i++ {
		s.Days = append(s.Days, NewWeather(s.Sun, i))
	}
}

/*
calculateRelevanDataFromDays Método para extraer los datos relevantes de los días calculados.
Se calcula:
Días de lluvia,
Días de sequía,
Días óptimos,
Perímetro máximo,
Picos de lluvia (Días en los cuales los planetas forman un triángulo con ese perímetro máximo) y
Días normales.
*/
func (s *SolarSystem) calculateRelevanDataFromDays() {
	//Se filtran los días por su condición climática.
	rainyDays 	:= s.filterByWeatherCondition(vo.Rain)
	dryDays 	:= s.filterByWeatherCondition(vo.Dry)
	optimalDays	:= s.filterByWeatherCondition(vo.Optimal)

	if rainyDays != nil {
		//Se obtiene el perímetro máximo de todos los días de lluvia.
		s.MaxPerimeter = findMaxPerimeter(rainyDays)

		//Se filtran los días de lluvia que tengan ese valor de perímetro máximo.
		//Luego, se extrae solo el número de día.
		s.DaysWithMaxRain = getDaysWithMaxPerimeter(rainyDays, s.MaxPerimeter)

		//Se asignan el total.
		s.RainyDays	= len(rainyDays)
	}

	if dryDays != nil {
		//Se asignan el total.
		s.DryDays = len(dryDays)
	}

	if optimalDays != nil {
		//Se asignan el total.
		s.OptimalDays = len(optimalDays)
	}

	s.NormalDays = len(s.Days) - (s.RainyDays + s.DryDays + s.OptimalDays)
}

//filterByWeatherCondition Método para filtrar los días por su condición climática valueobjects.WeatherCondition.
func (s *SolarSystem) filterByWeatherCondition(condition vo.WeatherCondition) (ret []*Weather) {
	for _, o := range s.Days {
		if o.WeatherCondition == condition {
			ret = append(ret, o)
		}
	}
	return
}

//findMaxPerimeter Método para encontrar el perímetro máximo de una lista de Weather.
func findMaxPerimeter(a []*Weather) (max float64) {
	max = a[0].Perimeter
	for _, o := range a {
		if o.Perimeter > max {
			max = o.Perimeter
		}
	}
	return
}

//getDaysWithMaxPerimeter Método para obtener la lista de días que tienen el perímetro máximo
func getDaysWithMaxPerimeter(a []*Weather, p float64) (ret []int) {
	for _, o := range a {
		if o.Perimeter == p {
			o.WeatherCondition = vo.RainPeak
			ret = append(ret, o.Day)
		}
	}
	return
}
