package valueobjects

//Weather Clase encargada del encapsulamiento de los datos correspondientes a un día específico en la predicción.
type Weather struct {
	Betasoide Star
	Vulcano   Star
	Ferengi   Star

	Perimeter        int
	WeatherCondition WeatherCondition
}
