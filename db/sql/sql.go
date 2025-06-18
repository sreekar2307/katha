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
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, fmt.Errorf("RDS Open gorm: %w", err)
	}
	db = db.Debug()

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("RDS Ping: %w", err)
	}
	return db.Omit(clause.Associations), nil
}
