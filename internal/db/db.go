package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pauleibye/urban-octo-bassoon/internal/entity"
)

func CreatePoint(db *sqlx.DB, point entity.Point) error {
	_, err := db.NamedExec(`INSERT INTO point (id, metric_id, time, value) VALUES (:id, :metric_id, :time, :value)`,
		point)
	return err
}

func CreateSeries(db *sqlx.DB, series entity.Series) error {
	_, err := db.NamedExec(`INSERT INTO series (id, name) VALUES (:id, :name)`,
		series)
	return err
}

func GetSeries(db *sqlx.DB, seriesId int) (entity.Series, error) {
	var series entity.Series
	err := db.Get(&series, `SELECT * FROM series WHERE id=$1 limit 1`, seriesId)
	return series, err
}

func GetPointsOfSeries(db *sqlx.DB, seriesId int) ([]entity.Point, error) {
	var points []entity.Point
	err := db.Select(&points, `SELECT * FROM point WHERE series_id=$1`, seriesId)
	if err != nil {
		return nil, err
	}
	return points, nil
}
