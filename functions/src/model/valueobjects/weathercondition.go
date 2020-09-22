package valueobjects

//WeatherCondition Enumerador de las posibles condiciones climáticas.
type WeatherCondition string

const (
	Rain     WeatherCondition = "lluvia"
	Normal   WeatherCondition = "normal"
	Dry      WeatherCondition = "sequia"
	RainPeak WeatherCondition = "pico"
	Optimal  WeatherCondition = "optima"
)
