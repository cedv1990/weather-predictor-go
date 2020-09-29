package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	"github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

type SolarSystemValidator struct {
	shareddomain.DomainModelValidator

	errors *[]shareddomain.Error
}

func NewSolarSystemValidator() *SolarSystemValidator {
	ss := new(SolarSystemValidator)
	ss.errors = &[]shareddomain.Error{}
	return ss
}

func (ss *SolarSystemValidator) Validate(model *model.SolarSystem) bool {
	return ss.errors == nil || len(*ss.errors) == 0
}

func (ss *SolarSystemValidator) GetErrors() *[]shareddomain.Error {
	return ss.errors
}