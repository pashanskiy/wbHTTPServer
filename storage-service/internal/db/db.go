package db

import (
	"errors"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	envParse "wbHTTPServer/storage-service/internal/env-parse"
)

type DB struct {
	Gorm *gorm.DB

	logger zerolog.Logger

	Config *Config
}

func NewDB(logger zerolog.Logger) *DB {
	db := &DB{
		logger: logger.
			With().
			Str("Logger", "DB").
			Logger(),
		Config: &Config{},
	}
	envParse.ReadEnvConfig(db.logger, db.Config)

	db.connectDB()
	db.migrateDB()

	return db
}

func (r *DB) connectDB() {

	r.createIfNotExist()

	r.logger.Info().Msg("Connecting to database...")

	gormLogLevel := logger.Warn
	if r.logger.Debug().Enabled() {
		gormLogLevel = logger.Info
	}

	gormDB, err := gorm.Open(
		sqlite.Open(r.Config.DataBaseFilePath),
		&gorm.Config{
			Logger: logger.Default.LogMode(gormLogLevel),
		})

	if err != nil {
		r.logger.Panic().Err(err).Msg("")
	}

	var version string
	if res := gormDB.Raw("SELECT sqlite_version();").First(&version); res.Error != nil {
		r.logger.Panic().Err(res.Error).Msg("")
	}

	r.logger.Info().Msgf("Database version: %s", version)

	r.Gorm = gormDB
}

func (r *DB) createIfNotExist() {
	var err error
	if _, err = os.Stat(r.Config.DataBaseFilePath); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(r.Config.DataBaseFilePath)
		r.logger.Info().Msgf("sqlite database has been successfully created. path: %s", r.Config.DataBaseFilePath)

	}

	if err != nil {
		r.logger.Fatal().Err(err).Msg("error create database")
	}
}

func (r *DB) migrateDB() {
	migrationsDirPath := r.Config.MigrationsDirPath
	log := r.logger

	if len(migrationsDirPath) == 0 {
		log.Warn().Msg("Directory with database migrations not specified - skip database migration")

		return
	}

	log.Info().Msg("Connecting to database for migration...")

	// Checking the availability of the directory with database migrations
	_, err := os.Stat(migrationsDirPath)
	if errors.Is(err, os.ErrNotExist) {
		log.Error().Err(err).Msgf("Directory with database migrations not found: %s", migrationsDirPath)

		return
	} else {
		log.Info().Msgf("Directory with database migrations: %s", migrationsDirPath)
	}

	db, err := r.Gorm.DB()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to connect to database to run migrations")
	}

	goose.SetDialect("sqlite3")
	if err = goose.Up(db, migrationsDirPath, goose.WithAllowMissing()); err != nil {
		log.Panic().Err(err).Msg("")
	}

	log.Info().Msg("Database migration completed")
}
