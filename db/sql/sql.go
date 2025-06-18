package sql

import (
	"database/sql"
	"fmt"
	"github.com/sreekar2307/khata/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var numberOfRetries = 3

func NewSqlConnection(sqlSettings config.SqlSettings) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d default_query_exec_mode=simple_protocol",
		sqlSettings.Host,
		sqlSettings.UserName,
		sqlSettings.Password,
		sqlSettings.DbName,
		sqlSettings.Port,
	)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("RDS Open: %w", err)
	}
	sqlDB.SetMaxOpenConns(sqlSettings.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(sqlSettings.MaxIdleConnections)
	sqlDB.SetConnMaxIdleTime(time.Duration(sqlSettings.MaxIdleConnectionTime) * time.Minute)
	var (
		db        *gorm.DB
		dbOpenErr error
	)
	for range numberOfRetries {
		db, dbOpenErr = gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if dbOpenErr != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	if dbOpenErr != nil {
		return nil, fmt.Errorf("RDS Open: %w", dbOpenErr)
	}
	db = db.Debug()

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("RDS Ping: %w", err)
	}
	return db.Omit(clause.Associations), nil
}
