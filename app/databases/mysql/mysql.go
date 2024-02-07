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
		log.Fatalf("failed to load MySQL configuration: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			config.MYSQL.MYSQL_USER,
			config.MYSQL.MYSQL_PASS,
			config.MYSQL.MYSQL_HOST,
			config.MYSQL.MYSQL_PORT,
			config.MYSQL.MYSQL_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil { 
		log.Fatalf("failed to connect to MySQL: %v", err)
	}

	migration.Migrate(db)

	log.Println("connected to MySQL")

	return db
}
