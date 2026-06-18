package database

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func Connect() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SSLMODE"),
	)
	return orm.RegisterDataBase("default", "postgres", connStr)
}
