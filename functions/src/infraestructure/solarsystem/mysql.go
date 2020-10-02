package solarsystem

import (
	"database/sql"
	"fmt"
	"github.com/cedv1990/weather-predictor-go/functions/src/model"
	"github.com/cedv1990/weather-predictor-go/functions/src/model/valueobjects"
	errors "github.com/cedv1990/weather-predictor-go/functions/src/shareddomain"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	dataBaseName     = "weather_predictions"
	dataBaseUser     = "root"
	dataBaseUserPass = "123"
	dataBasePort     = 3306
)

type MySqlRepository struct {
	RepositoryBase
}

func NewMySqlRepository() *MySqlRepository {
	iss := new(MySqlRepository)
	return iss
}

func (iis *MySqlRepository) Save(solarSystem *model.SolarSystem) (*model.SolarSystem, *errors.ValidationException) {
	err := prepare()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	for _, w := range solarSystem.Days {
		st, err := insert(SQLINSERTSOLARSYSTEM)
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

func prepare() *errors.ValidationException {
	conn, err := openConnectionMySql()
	if err != nil {
		ex := errors.NewValidationException(&[]errors.Error{errors.NewMySqlError(errors.WithoutConnection)})
		return ex
	}
	db = conn

	_, err = createTable(SQLWEATHER)
	if err != nil {
		ex := errors.NewValidationException(&[]errors.Error{errors.NewMySqlError(errors.TableNotCreated)})
		return ex
	}

	return nil
}

func openConnectionMySql() (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/", dataBaseUser, dataBaseUserPass, dataBasePort))
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec("CREATE DATABASE IF NOT EXISTS " + dataBaseName)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec("USE " + dataBaseName)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createTable(sql string) (sql.Result, error) {
	return db.Exec(sql)
}

func insert(sql string) (*sql.Stmt, error) {
	return db.Prepare(sql)
}

func queryRow(query string, pointer interface{}) (*sql.Row, error) {
	a := db.QueryRow(query)
	err := a.Scan(pointer)

	return a, err
}