package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-assignment3/internal/entity"
)

type weatherRepositoryImpl struct {
	DB *sql.DB
}

type WeatherRepository interface {
	UpdateValue(windValue string, waterValue string, report string) error
	GetReport(ctx context.Context) (entity.StatusReport, error)
}

func ProvideWeatherRepository(DB *sql.DB) *weatherRepositoryImpl {
	return &weatherRepositoryImpl{DB: DB}
}

var (
	UPDATE_DATA     = "UPDATE status SET wind = ?, water = ?, report = ? WHERE id = 1"
	GET_REPORT_DATA = "SELECT wind, water, report FROM status"
)

func (w weatherRepositoryImpl) UpdateValue(windValue string, waterValue string, report string) error {
	query := UPDATE_DATA

	stmt, err := w.DB.Prepare(query)
	if err != nil {
		log.Printf("[UpdateValue] failed to prepare the statement, err => %v", err)
		return err
	}
	_, err = stmt.Exec(windValue, waterValue, report)
	if err != nil {
		log.Printf("[UpdateValue] failed to execute to the database")
		return err
	}
	return nil
}

func (w weatherRepositoryImpl) GetReport(ctx context.Context) (entity.StatusReport, error) {
	query := GET_REPORT_DATA

	stmt, err := w.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetReport] failed to prepare the statement")
		return entity.StatusReport{}, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[GetReport] failed to query to the database")
		return entity.StatusReport{}, err
	}

	report := entity.StatusReport{}
	for rows.Next() {
		details := entity.StatusReport{}
		err := rows.Scan(
			&details.Weather.Wind,
			&details.Weather.Water,
			&details.Status,
		)
		if err != nil {
			log.Printf("[GetReport] failed to scan the data")
			return entity.StatusReport{}, err
		}
		report = details
	}
	return report, nil
}
