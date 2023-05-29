package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/i-akbarshoh/auth-service/internal/pkg/config"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"os"
)

func NewDB() *bun.DB {
	//dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.C.Database.User, config.C.Database.Password, config.C.Database.Host, config.C.Database.Port, config.C.Database.DBName)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.C.Database.Host, config.C.Database.Port, config.C.Database.User, config.C.Database.Password, config.C.Database.DBName)
	pDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := pDB.Ping(); err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(pDB, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	migrateUp(db.DB)
	return db
}

func migrateUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/pkg/script/migration",
		config.C.Database.DBName, driver)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println(err)
		os.Exit(1)
	}
}
