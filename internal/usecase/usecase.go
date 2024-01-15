package usecase

import "github.com/jmoiron/sqlx"

type Usecase struct {
	DB *sqlx.DB
}
