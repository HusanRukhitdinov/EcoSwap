package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectionDb() (*sql.DB, error) {
	// cnf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"localhost", 5432, "postgres", "users_service", "3333")
	return sql.Open("postgres", conDb)
}
