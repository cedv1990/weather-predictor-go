package weather

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/http/controllers/solarsystem"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	weathercondition "github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	"github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var controller *WeatherController

func TestMain(m *testing.M) {
	solarController := solarsystem.NewSolarSystemController()
	solarController.Generate(3652)

	controller = NewWeatherController()
	code := m.Run()
	//Some code here
	os.Exit(code)
}

func Test_getDayWeatherFromSolarSystem_ByZeroMustBeNormal(t *testing.T) {
	ass := assert.New(t)

	const day = 0

	weatherExpectedByZero := &model.Weather{
		Day: day,
		Perimeter: 0,
		WeatherCondition: weathercondition.Dry,
	}

	controller.getDayWeatherFromSolarSystem(day)

	errors := controller.GetErrors()
	weather := controller.data

	NoErrorValidation(*ass, errors)

	CompareValues(ass, "Test_getDayWeatherFromSolarSystem_ByZeroMustBeNormal", weatherExpectedByZero, weather)
}

func Test_getDayWeatherFromSolarSystem_By566MustBeRain(t *testing.T) {
	ass := assert.New(t)

	const day = 566

	weatherExpected := &model.Weather{
		Day: day,
		Perimeter: 6129.256444994509,
		WeatherCondition: weathercondition.Rain,
	}

	controller.getDayWeatherFromSolarSystem(day)

	errors := controller.GetErrors()
	weather := controller.data

	NoErrorValidation(*ass, errors)

	CompareValues(ass, "Test_getDayWeatherFromSolarSystem_By566MustBeRain", weatherExpected, weather)
}

func Test_getDayWeatherFromSolarSystem_By4000(t *testing.T) {
	ass := assert.New(t)

	const (
		day = 4000
		testName = "Test_getDayWeatherFromSolarSystem_By4000"
	)

	controller.getDayWeatherFromSolarSystem(day)

	errors := controller.GetErrors()
	weather := controller.data

	ass.Nil(weather, "%s: Weather should be nil.", testName)
	ass.NotNil(errors, "%s: Errors should exists.", testName)

	error := (*errors)[0]

	ass.EqualValuesf(shareddomain.NotExistsErrorName, error.GetMessage(), "%s: The error should be %s", testName, shareddomain.NotExistsErrorName)
}

func Test_getDayWeatherFromSolarSystem_ByNegative(t *testing.T) {
	ass := assert.New(t)

	const (
		day = -1
		testName = "Test_getDayWeatherFromSolarSystem_ByNegative"
	)

	controller.getDayWeatherFromSolarSystem(day)

	errors := controller.GetErrors()
	weather := controller.data

	ass.Nil(weather, "%s: Weather should be nil.", testName)
	ass.NotNil(errors, "%s: Errors should exists.", testName)
}

func NoErrorValidation(ass assert.Assertions, errors *[]shareddomain.Error) {
	ass.Nil(errors, "It should be no errors.")
}

func CompareValues(ass *assert.Assertions, testFunction string, weatherExpected, weatherQueried *model.Weather) {
	ass.EqualValuesf(weatherExpected.WeatherCondition, weatherQueried.WeatherCondition, "%s: Weather condition should be %s", testFunction, weatherExpected.WeatherCondition)
	ass.EqualValuesf(weatherExpected.Day, weatherQueried.Day, "%s: Day number should be %s", testFunction, weatherExpected.Day)
	ass.EqualValuesf(weatherExpected.Perimeter, weatherQueried.Perimeter, "%s: Perimeter should be %f", testFunction, weatherExpected.Perimeter)
}