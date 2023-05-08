package db

import (
	"database/sql"
	"embed"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/mauroccvieira/restapi-echo-go/logger"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func New(development bool) (*sql.DB, error) {
	var uri string
	if development {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	} else {
		uri = "postgres://postgres:postgres@echo-db:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationsFS,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return nil, err
	}

	// Run all seed sql file if on dev
	if development {
		if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS seeds (
			id SERIAL PRIMARY KEY,
			filename TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		);`); err != nil {
			return nil, err
		}

		seedFiles, err := os.ReadDir("db/seed/")
		if err != nil {
			logger.Error("failed to read seed files", zap.Error(err))
		}

		executedSeeds := make(map[string]bool)
		rows, err := db.Query("SELECT filename FROM seeds")
		if err != nil {
			logger.Error("failed to query seeds table", zap.Error(err))
		}
		defer rows.Close()

		for rows.Next() {
			var filename string
			if err := rows.Scan(&filename); err != nil {
				logger.Error("failed to scan seeds table row", zap.Error(err))
			}
			executedSeeds[filename] = true
		}

		for _, f := range seedFiles {
			filename := f.Name()
			if executedSeeds[filename] {
				continue
			}

			c, err := os.ReadFile("db/seed/" + filename)
			if err != nil {
				logger.Error("failed to read seed file", zap.Error(err))
			}

			sqlCode := string(c)

			_, err = db.Exec(sqlCode)
			if err != nil {
				logger.Error("failed to seed database", zap.Error(err))
			}

			if _, err := db.Exec("INSERT INTO seeds (filename) VALUES ($1)", filename); err != nil {
				logger.Error("failed to insert row into seeds table", zap.Error(err))
			}
		}

		logger.Info("seeded database")
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Mock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		logger.Fatal("failed to create mock db", zap.Error(err))
	}

	return db, mock
}
