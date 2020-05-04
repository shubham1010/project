package common

import (
	"os"
	"database/sql"

	"github.com/shubham1010/project/temp/migrations"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
)

type DBEnv struct {
	Server, DBUser, DBPass, DBHost, DBName, Port string
}

var Env DBEnv

func initConfig() {
	LoadEnv()
	setEnvVariables()
}

func setEnvVariables() {
	Env.Server = getEnvAvaiableValue("SERVER")
	Env.DBUser = getEnvAvaiableValue("MYSQL_USER")
	Env.DBPass = getEnvAvaiableValue("MYSQL_PASSWORD")
	Env.DBHost = getEnvAvaiableValue("MYSQL_HOST")
	Env.DBName = getEnvAvaiableValue("DATABASE")
	Env.Port = getEnvAvaiableValue("PORT")


	if (Env.DBUser==" " || Env.DBPass==" " || Env.DBHost==" " || Env.DBName==" " || Env.Port==" ") {
		log.Warn("[COMMON]: Env Variable is not set")
	}

}

func LoadEnv() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(dir+"/.env")
	if err != nil {
	  log.Fatalf("[COMMON]: Error loading .env file")
	}

}

func getEnvAvaiableValue(key string) string {
	return os.Getenv(key)
}

var db *sql.DB

func GetDBConnection() *sql.DB {
	if db == nil {
		var err error

		db, err = sql.Open("mysql",Env.DBUser+":"+Env.DBPass+"@tcp("+Env.DBHost+":"+Env.Port+")/"+Env.DBName)

		if err != nil {
			log.Fatalf("[COMMOM]: GetDBConnection%s\n", err)
		}

		if err := migrations.MigrateDatabase(db); err!=nil {
			log.Fatalf("[COMMON]: GetDBConnection migrate database %s\n", err)
		}

	}
	return db
}

func createDBConnection() {
	var err error

	db, err = sql.Open("mysql",Env.DBUser+":"+Env.DBPass+"@tcp("+Env.DBHost+":"+Env.Port+")/"+Env.DBName)

	if err != nil {
		log.Fatalf("[COMMON]: CreateDBConnection %s\n", err)
	}

	if err := migrations.MigrateDatabase(db); err!=nil {
			log.Fatalf("[COMMON]: CreateDBConnection Migrate database %s\n", err)
	}
}
