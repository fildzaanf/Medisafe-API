package mysql

import (
	"fmt"
	"log"
	"talkspace/app/configs"
	"talkspace/app/databases/migration"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() *gorm.DB {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load mysql configuration: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			config.MYSQL.DB_USER,
			config.MYSQL.DB_PASS,
			config.MYSQL.DB_HOST,
			config.MYSQL.DB_PORT,
			config.MYSQL.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}

	migration.Migrate(db)

	log.Println("connected to mysql")

	return db
}
