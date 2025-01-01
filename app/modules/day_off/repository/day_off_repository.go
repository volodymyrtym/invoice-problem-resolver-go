package repository

import (
	"context"
	"database/sql"
)

type DayOffRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewDayOffRepository(db *sql.DB, ctx context.Context) *DayOffRepository {
	return &DayOffRepository{db: db, ctx: ctx}
}
