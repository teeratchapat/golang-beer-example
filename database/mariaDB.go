package database

import (
	"context"
	"database/sql"
	"fmt"
	config "golang-beer-example/configs"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // Use this for MariaDB driver
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var maria *sql.DB

func InitialMariaDatabase() error {
	if maria != nil {
		_ = maria.Close()
		maria = nil
	}

	// connStr := config.AppConfig.DB.Maria.DSN
	conn, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/beer_db")
	if err != nil {
		log.Printf(" | InitialMariaDatabase Failed | MSG : %v", err)
		return err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(10 * time.Second)
	conn.SetConnMaxIdleTime(10 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = conn.PingContext(ctx); err != nil {
		log.Printf(" | InitialMariaDatabase Ping Failed | MSG : %v", err)
		conn.Close()
		return err
	}

	maria = conn
	log.Println(" | InitialMariaDatabase Success ")
	return nil
}

func ConnectMariaDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DB.Maria.Username,
		config.AppConfig.DB.Maria.Password,
		config.AppConfig.DB.Maria.Host,
		config.AppConfig.DB.Maria.Port,
		config.AppConfig.DB.Maria.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MariaDB using GORM: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get *sql.DB from GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	sqlDB.SetConnMaxIdleTime(10 * time.Second)

	log.Println(" | ConnectMariaDB Success ")
	return db
}
