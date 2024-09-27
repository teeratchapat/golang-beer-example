package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Build struct {
		Date   string
		Number int
	}
	Service struct {
		Name         string
		Port         string
		Endpoint     string
		InfoEndpoint string
	}
	DB struct {
		MongoDB struct {
			URI    string
			DBName string
		}
		Maria struct {
			DriverName string
			DSN        string
			Username   string
			Password   string
			Host       string
			Port       string
			Database   string
		}
	}
	Volumes struct {
		MongoData string
	}
}

var AppConfig *Config

func LoadConfig() {

	a := os.Args
	env := "local"
	if len(a) > 1 {
		env = os.Args[1]
	}
	path := "../"
	if env != "local" {
		path = "./"
	}

	viper.SetConfigFile(path + ".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	viper.SetDefault("BUILD_DATE", "unknown")
	viper.SetDefault("BUILD_NUMBER", 1)

	AppConfig = &Config{
		Build: struct {
			Date   string
			Number int
		}{
			Date:   viper.GetString("BUILD_DATE"),
			Number: viper.GetInt("BUILD_NUMBER"),
		},
		Service: struct {
			Name         string
			Port         string
			Endpoint     string
			InfoEndpoint string
		}{
			Name:         viper.GetString("SERVICE_NAME"),
			Port:         viper.GetString("SERVICE_PORT"),
			Endpoint:     viper.GetString("SERVICE_ENDPOINT"),
			InfoEndpoint: viper.GetString("SERVICE_INFO_ENDPOINT"),
		},
		DB: struct {
			MongoDB struct {
				URI    string
				DBName string
			}
			Maria struct {
				DriverName string
				DSN        string
				Username   string
				Password   string
				Host       string
				Port       string
				Database   string
			}
		}{
			MongoDB: struct {
				URI    string
				DBName string
			}{
				URI:    viper.GetString("MONGODB_URI"),
				DBName: viper.GetString("MONGODB_DBNAME"),
			},
			Maria: struct {
				DriverName string
				DSN        string
				Username   string
				Password   string
				Host       string
				Port       string
				Database   string
			}{
				DriverName: viper.GetString("MARIA_DB_DRIVER_NAME"),
				DSN:        viper.GetString("MARIA_DB_DSN"),
				Username:   viper.GetString("DB_USER"),
				Password:   viper.GetString("DB_PASSWORD"),
				Host:       viper.GetString("DB_HOST"),
				Port:       viper.GetString("DB_PORT"),
				Database:   viper.GetString("DB_NAME"),
			},
		},
		Volumes: struct {
			MongoData string
		}{
			MongoData: viper.GetString("VOLUME_MONGO_DATA"),
		},
	}
}
