package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	exceptions "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

func Create(days int) (*model.SolarSystem, *exceptions.ValidationException) {
	solarSystemValidator := NewSolarSystemValidator()
	solarSystem := model.NewSolarSystem(days)

	if !solarSystemValidator.Validate(solarSystem) {
		errors := solarSystemValidator.GetErrors()
		ex := exceptions.NewValidationException(errors)
		return nil, ex
	}

	return solarSystem, nil
}
