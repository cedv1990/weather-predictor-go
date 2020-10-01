package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	"github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var controller *SolarSystemController

func TestMain(m *testing.M) {
	controller = NewSolarSystemController()
	code := m.Run()
	//Some code here
	os.Exit(code)
}

func Test_generate_WithOneDay(t *testing.T) {
	ass := assert.New(t)

	const days = 1

	solarExpectedOneDay := &model.SolarSystem{
		DaysWithMaxRain: []int{},
		MaxPerimeter: -1,
		DryDays: 1,
		RainyDays: 0,
		OptimalDays: 0,
		NormalDays: 0,
		Days: make([]*model.Weather, days),
	}

	controller.ResetRepo()
	controller.Generate(days)

	errors := controller.GetErrors()
	solar := controller.data

	NoErrorValidation(*ass, errors)

	CompareValues(ass, "Test_generate_WithOneDay", solarExpectedOneDay, solar)
}

func Test_generate_WithFiftyDays(t *testing.T) {
	ass := assert.New(t)

	const days = 50

	solarExpectedFiftyDays := &model.SolarSystem{
		Days: make([]*model.Weather, days),
		DaysWithMaxRain: make([]int, 1),
		MaxPerimeter: 6135.928957012673,
		DryDays: 1,
		NormalDays: 41,
		OptimalDays: 0,
		RainyDays: 8,
	}

	controller.ResetRepo()
	controller.Generate(days)

	errors := controller.GetErrors()
	solar := controller.data

	NoErrorValidation(*ass, errors)

	CompareValues(ass, "Test_generate_WithFiftyDays", solarExpectedFiftyDays, solar)
}

func Test_generate_WithNegativeDays(t *testing.T) {
	ass := assert.New(t)

	controller.ResetRepo()
	controller.Generate(-1)

	errors := controller.GetErrors()
	solar := controller.data

	ass.Nil(solar, "Test_generate_WithNegativeDays: SolarSystem should be nil.")
	ass.NotNil(errors, "Test_generate_WithNegativeDays: Errors should exists.")
}

func NoErrorValidation(ass assert.Assertions, errors *[]shareddomain.Error) {
	ass.Nil(errors, "It should be no errors.")
}

func CompareValues(ass *assert.Assertions, testFunction string, solarExpected, solarGenerated *model.SolarSystem) {
	ass.EqualValuesf(cap(solarExpected.Days), len(solarGenerated.Days), "%s: It should have %d days calculated", testFunction, cap(solarExpected.Days))
	ass.EqualValuesf(solarExpected.MaxPerimeter, solarGenerated.MaxPerimeter, "%s: MaxPerimeter should be %f", testFunction, solarExpected.MaxPerimeter)
	ass.EqualValuesf(solarExpected.DryDays, solarGenerated.DryDays, "%s: DryDays should be %d", testFunction, solarExpected.DryDays)
	ass.EqualValuesf(solarExpected.NormalDays, solarGenerated.NormalDays, "%s: NormalDays should be %d", testFunction, solarExpected.NormalDays)
	ass.EqualValuesf(solarExpected.OptimalDays, solarGenerated.OptimalDays, "%s: OptimalDays should be %d", testFunction, solarExpected.OptimalDays)
	ass.EqualValuesf(solarExpected.RainyDays, solarGenerated.RainyDays, "%s: RainyDays should be %d", testFunction, solarExpected.RainyDays)
	ass.EqualValuesf(cap(solarExpected.DaysWithMaxRain), len(solarGenerated.DaysWithMaxRain), "%s: RainyDays should be %d", testFunction, solarExpected.RainyDays)
}