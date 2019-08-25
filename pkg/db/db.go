package db

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/rphillips/escapepod/pkg/models"
)

type DBParams struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func openPostgres(dbURL string) (*gorm.DB, error) {
	if len(dbURL) == 0 {
		dbURL = strings.Join([]string{
			"host=" + viper.GetString("db.host"),
			"port=" + viper.GetString("db.port"),
			"user=" + viper.GetString("db.user"),
			"password=" + viper.GetString("db.password"),
			"dbname=" + viper.GetString("db.dbname"),
			"sslmode=" + viper.GetString("db.sslmode"),
		}, " ")
		sanitizedDBURL := strings.Join([]string{
			"host=" + viper.GetString("db.host"),
			"port=" + viper.GetString("db.port"),
			"user=" + "<****>",
			"password=" + "<****>",
			"dbname=" + viper.GetString("db.dbname"),
			"sslmode=" + viper.GetString("db.sslmode"),
		}, " ")
		log.Debug("db.url", sanitizedDBURL)
	}

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func openSqlite(dbURL string) (*gorm.DB, error) {
	return gorm.Open("sqlite3", dbURL)
}

func OpenDB() (*gorm.DB, error) {
	dbURL := viper.GetString("db.url")
	dbType := viper.GetString("db.type")

	var db *gorm.DB
	var err error

	switch dbType {
	case "postgres":
		db, err = openPostgres(dbURL)
	case "sqlite":
		db, err = openSqlite(dbURL)
	default:
		err = errors.New("invalid database")
	}

	if err != nil {
		return nil, err
	}

	return models.Migrate(db), nil
}

func NewString(s string) *string {
	p := new(string)
	*p = s
	return p
}
