package repository

import (
	"context"
	"fmt"
	"github.com/YuStep/wbl0/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

//type Config struct {
//	Host     string
//	Port     string
//	Username string
//	Password string
//	DBName   string
//	SSLMode  string
//}

//func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
//	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
//		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
//	if err != nil {
//		return nil, err
//	}
//
//	err = db.Ping()
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}

func Connect(cfg *config.PG) (*pgxpool.Pool, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	return pgxpool.Connect(context.Background(), connection)
}
