package entity

import "time"

type Point struct {
	Id       int       `db:"id"`
	MetricId int       `db:"series_id"`
	Time     time.Time `db:"time"`
	Value    float64   `db:"value"`
}

type Series struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
