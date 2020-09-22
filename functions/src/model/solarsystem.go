package model

//SolarSystem Entidad encargada del encapsulamiento de las propiedades de cada estrella.
type SolarSystem struct {
}

type Weather struct {
}

type WeatherCondition string

const (
	Rain     WeatherCondition = "lluvia"
	Normal   WeatherCondition = "normal"
	Dry      WeatherCondition = "sequia"
	RainPeak WeatherCondition = "pico"
	Optimal  WeatherCondition = "optima"
)
