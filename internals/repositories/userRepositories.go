package repositories

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (h *UserRepository) Create() {

}
