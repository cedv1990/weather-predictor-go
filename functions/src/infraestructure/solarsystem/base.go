package solarsystem

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

type RepositoryBase struct {
	domain.Repository
}

func (iis *RepositoryBase) SendError() (*model.Weather, *errors.ValidationException) {
	ex := errors.NewValidationException(&[]errors.Error{errors.NewNotExistsError(true)})
	return nil, ex
}