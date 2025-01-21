package postgresql

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(
	dsn string,

) (*sql.DB, *gorm.DB, error) {
	if len(dsn) == 0 {
		return nil, nil, fmt.Errorf("sources database connetions is empty")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "wa.",
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	connection, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	connection.SetConnMaxLifetime(time.Minute)
	connection.SetMaxIdleConns(15)
	connection.SetMaxOpenConns(15)

	fmt.Println("db connected")

	return connection, db, nil
}
