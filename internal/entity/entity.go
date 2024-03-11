package entity

import "time"

type Point struct {
	Id       int       `db:"id"`
	MetricId int       `db:"series_id"`
	Time     time.Time `db:"time"`
	Value    float64   `db:"value"`
}

type ChartSeries struct {
	ChartId  int `db:"chart_id"`
	SeriesId int `db:"series_id"`
}

type Series struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type Chart struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
