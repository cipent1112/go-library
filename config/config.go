package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-library/config/migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	AppConfig struct {
		Name string
		Port int
	}
	DBConfig struct {
		Engine    string
		Host      string
		Port      int
		Username  string
		Password  string
		Schema    string
		DebugMode bool
		Orm       *gorm.DB
	}
}

var Conf Config

func init() {
	loadEnvVars()
	loadAppConfig()
	connectDB()
	migrate.AutoMigrate(Conf.DBConfig.Orm)
}

/* Load env.json */
func loadEnvVars() {
	viper.SetEnvPrefix("go-library")
	errBind := viper.BindEnv("env")

	if errBind != nil {
		panic(fmt.Errorf(errBind.Error()))
	}

	currentDirectory, _ := os.Getwd()
	viper.AddConfigPath(fmt.Sprintf("%s/config/", currentDirectory))
	viper.SetConfigName("env.json")
	viper.SetConfigType("json")

	errRead := viper.ReadInConfig()

	if errRead != nil {
		panic(fmt.Errorf("fatal error config file: %s", errRead.Error()))
	}
}

func loadAppConfig() {
	appConfig := Conf.AppConfig
	appConfig.Name = viper.GetString("app.name")
	appConfig.Port = viper.GetInt("app.port")
	Conf.AppConfig = appConfig
}

func connectDB() {
	dbConfig := Conf.DBConfig
	dbConfig.Engine = viper.GetString("database.engine")
	dbConfig.Host = viper.GetString("database.host")
	dbConfig.Port = viper.GetInt("database.port")
	dbConfig.Username = viper.GetString("database.username")
	dbConfig.Password = viper.GetString("database.password")
	dbConfig.Schema = viper.GetString("database.schema")
	dbConfig.DebugMode = viper.GetBool("database.debug_mode")

	conf := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Schema,
	)

	db, err := gorm.Open(
		mysql.Open(conf),
	)
	if err != nil {
		panic("failed to connect to database")
	}

	if dbConfig.DebugMode {
		dbConfig.Orm = db.Debug()
		return
	}

	dbConfig.Orm = db
	Conf.DBConfig = dbConfig
	log.Println("Success connect to mysql database")
}
