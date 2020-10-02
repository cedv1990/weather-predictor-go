package solarsystem

import (
	"database/sql"
	"fmt"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	"github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //Propiedad que tendrá la conexión a MySQL

const (
	dataBaseName     = "weather_predictions"
	dataBaseUser     = "root"
	dataBaseUserPass = "123"
	dataBasePort     = 3306
)

//MySqlRepository Clase encargada de la persistencia de datos en MySQL.
type MySqlRepository struct {
	RepositoryBase
}

func NewMySqlRepository() *MySqlRepository {
	iss := new(MySqlRepository)
	return iss
}

//Save Método para almacenar los datos generados en MySQL.
func (iis *MySqlRepository) Save(solarSystem *model.SolarSystem) (*model.SolarSystem, *errors.ValidationException) {
	err := prepare()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	for _, w := range solarSystem.Days {
		st, err := db.Prepare(SQLINSERTSOLARSYSTEM)
		if err != nil {
			ex := errors.NewValidationException(&[]errors.Error{errors.NewMySqlError(errors.NotInserted)})
			return nil, ex
		}

		_, er := st.Exec(w.Day, w.WeatherCondition)

		if er != nil {
			ex := errors.NewValidationException(&[]errors.Error{errors.NewMySqlError(errors.NotInserted)})
			return nil, ex
		}
	}

	return solarSystem, nil
}

//Exists Método para validar si ya existen datos en memoria.
func (iis *MySqlRepository) Exists() (error bool) {
	er := prepare()
	defer db.Close()

	if er != nil {
		return false
	}

	_, err := queryRow(SQLEXISTS, &error)
	if err != nil {
		error = true
	}
	return
}

//GetDay Método para obtener el estado del clima de un día específico.
func (iis *MySqlRepository) GetDay(day int) (*model.Weather, *errors.ValidationException) {
	if !iis.Exists() {
		return iis.SendError()
	}

	er := prepare()
	defer db.Close()

	if er != nil {
		return iis.SendError()
	}

	var weatherCondition string

	_, err := queryRow(fmt.Sprintf(SQLQUERY, day), &weatherCondition)
	if err != nil {
		return iis.SendError()
	}

	return &model.Weather{
		Day: day,
		WeatherCondition: valueobjects.WeatherCondition(weatherCondition),
	}, nil
}

//prepare Método para crear la tabla si no existe.
func prepare() *errors.ValidationException {
	conn, err := openConnectionMySql()
	if err != nil {
		ex := errors.NewValidationException(&[]errors.Error{errors.NewMySqlError(errors.WithoutConnection)})
		return ex
	}
	db = conn

	_, err = db.Exec(SQLWEATHER)
	if err != nil {
		ex := errors.NewValidationException(&[]errors.Error{errors.NewMySqlError(errors.TableNotCreated)})
		return ex
	}

	return nil
}

//openConnectionMySql Método para crear la base de datos si no existe y conectarse a ella.
func openConnectionMySql() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/", dataBaseUser, dataBaseUserPass, dataBasePort))
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dataBaseName)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("USE " + dataBaseName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

//queryRow Método para consultar el dato de una fila.
func queryRow(query string, pointer interface{}) (*sql.Row, error) {
	a := db.QueryRow(query)
	err := a.Scan(pointer)

	return a, err
}