package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// auto migrate the following models.
var modelsToMigrate = []interface{}{
	&User{},
	&Referral{},
}

// InitDB initializes the DB.
func InitDB(username, password, host, port, dbName string) error {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	log.Println(dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return err
	}
	db = gormDB
	migrateModels()
	return nil
}

// CloseDB closes the underlying database connection.
func CloseDB() {
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.Close()
	}
}

func ClearTables(models ...any) {
	for _, model := range models {
		switch model.(type) {
		case *User:
			db.Exec("truncate table `users`")
		case *Referral:
			db.Exec("truncate table `referrals`")
		}
	}
}

func migrateModels() {
	if err := db.AutoMigrate(modelsToMigrate...); err != nil {
		log.WithError(err).
			Fatal("failed to migrate models")
	}
}
