package controller

import "github.com/uptrace/bun"

type HealthDB struct {
	db *bun.DB
}

func NewHealthDB(db *bun.DB) *HealthDB {
	return &HealthDB{
		db: db,
	}
}

func (h HealthDB) IsActive() bool {
	if err := h.db.Ping(); err != nil {
		return false
	}
	return true
}
