package datebase

import (
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	sqlite "github.com/ytsruh/gorm-libsql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	dbToken := os.Getenv("TURSO_TOKEN")
	env := os.Getenv("ENV")
	fullURLKey := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)
	var db *gorm.DB
	if env == "local" {
		db, err := gorm.Open(sqlite.Open("my-health.db"), &gorm.Config{})
		db = db.Debug()
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DSN:        fullURLKey,
		DriverName: "libsql",
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
