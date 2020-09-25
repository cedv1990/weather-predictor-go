package solarsystem

import (
	domain "github.com/cedv1990/weather-predictor-go/functions/src/domain/solarsystem"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

var (
	repo *model.SolarSystem
)

type InMemorySolarSystemRepository struct {
	domain.SolarSystemRepository
}

func NewInMemorySolarSystemRepository() *InMemorySolarSystemRepository {
	iss := new(InMemorySolarSystemRepository)
	repo = nil
	return iss
}

func (iss InMemorySolarSystemRepository) Create(solarSystem *model.SolarSystem) (*model.SolarSystem, *errors.ValidationException) {
	if iss.Exists() {
		ex := errors.NewValidationException([]errors.Error{errors.AlreadyExistsError{ Is: true }})
		return nil, ex
	}
	repo = solarSystem
	return solarSystem, nil
}

func (iss InMemorySolarSystemRepository) Exists() bool {
	return repo != nil
}

func (iss InMemorySolarSystemRepository) GetDay(day int) (*model.Weather, *errors.ValidationException) {

	sendError := func() (*model.Weather, *errors.ValidationException) {
		ex := errors.NewValidationException([]errors.Error{errors.NotExistsError{ No: true }})
		return nil, ex
	}

	if iss.Exists() {
		return sendError()
	}

	weather := repo.Days[day]

	if weather == nil {
		return sendError()
	}

	return weather, nil
}