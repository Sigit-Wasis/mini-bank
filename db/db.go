package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // or the appropriate driver for your DB_DRIVER
)

type Config struct {
    DBDriver      string
    DBSource      string
    ServerAddress string
}

// LoadConfig membaca variabel lingkungan dari file .env
func LoadConfig(path string) (Config, error) {
    var err error

    // Load file .env jika ada (opsional di production)
    err = godotenv.Load(path)
    if err != nil {
        return Config{}, err
    }

    config := Config{
        DBDriver:      os.Getenv("DB_DRIVER"),
        DBSource:      os.Getenv("DB_SOURCE"),
        ServerAddress: os.Getenv("SERVER_ADDRESS"),
    }

    // Validasi sederhana
    if config.DBDriver == "" || config.DBSource == "" {
        return Config{}, fmt.Errorf("environment variables DB_DRIVER dan DB_SOURCE wajib diisi")
    }

    return config, nil
}

// NewDB membuka koneksi database
func NewDB(config Config) (*sql.DB, error) {
    db, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}