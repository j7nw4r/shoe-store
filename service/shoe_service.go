package service

import "database/sql"

type ShoeService struct {
	db *sql.DB
}

func NewShoeService(db *sql.DB) *ShoeService {
	return &ShoeService{
		db: db,
	}
}
