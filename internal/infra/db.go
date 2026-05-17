package infra

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
