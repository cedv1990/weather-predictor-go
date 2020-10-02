package solarsystem

import (
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
)

var repo *model.SolarSystem //Propiedad que tendrá los datos en memoria.

//InMemoryRepository Clase encargada de la persistencia de datos en memoria.
type InMemoryRepository struct {
	RepositoryBase
}

func NewInMemoryRepository() *InMemoryRepository {
	iss := new(InMemoryRepository)
	repo = nil
	return iss
}

//Create Método para almacenar los datos generados en memoria.
func (iss *InMemoryRepository) Save(solarSystem *model.SolarSystem) (*model.SolarSystem, *errors.ValidationException) {
	//Se valida si los datos ya fueron generados.
	if iss.Exists() {
		ex := errors.NewValidationException(&[]errors.Error{errors.NewAlreadyExistsError(true)})
		return nil, ex
	}
	repo = solarSystem
	return solarSystem, nil
}

//Exists Método para validar si ya existen datos en memoria.
func (iss *InMemoryRepository) Exists() bool {
	return repo != nil
}

//GetDay Método para obtener el estado del clima de un día específico.
func (iss *InMemoryRepository) GetDay(day int) (*model.Weather, *errors.ValidationException) {

	//Se valida si hay datos.
	if !iss.Exists() {
		return iss.SendError()
	}

	//Se valida si el día existe
	if day < len(repo.Days) && day >= 0 {
		//Se filtra por el número de día.
		weather := repo.Days[day]

		//Se valida si el día tiene datos.
		if weather == nil {
			return iss.SendError()
		}

		return weather, nil
	}

	return iss.SendError()
}