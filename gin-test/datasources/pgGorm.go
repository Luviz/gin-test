package datasources

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDsnConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

func (c PostgresDsnConfig) DSN() string {
	return PGDSN(c.Host, c.User, c.Password, c.Database, c.Port)
}

func PGDSN(host, user, password, dbName, port string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, port)
}

func InitGORM(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
