package db

type Config struct {
	DataBaseFilePath  string `env:"DB_FILE" envDefault:"./storage-service/sqlite.db"`
	MigrationsDirPath string `env:"DB_MIGRATIONS_DIR" envDefault:"./storage-service/internal/db/migrations"`
}
