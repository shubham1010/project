package common

import (
	"net/http"
	"encoding/json"
	"log"
	"os"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

type MigrationLogger struct {
	verbose bool
}

type (
	appError struct {
		Error		string `json: "error"`
		Message		string `json: "message"`
		HttpStatus	int `json: "status"`
	}

	errorResource struct {
		Data appError `json: "data"`
	}

	configuration struct {
		Server, Database, DBUser, DBPwd, MysqlDBHost string
	}
)

func (ml *MigrationLogger) Printf(format string, v ...interface{}) {
	log.Printf(format,v)
}

func (ml *MigrationLogger) Verbose() bool {
	return ml.verbose
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError {
		Error:		handlerError.Error(),
		Message:	message,
		HttpStatus:	code,
	}

	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Context-Type", "application/json ; charset=utf-8")
	w.WriteHeader(code)

	if j, err := json.Marshal(errorResource{Data: errObj}); err==nil {
		w.Write(j)
	}
}

var AppConfig configuration

func initConfig() {
	/*dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(dir+"/containerData/config.json")*/
	file, err := os.Open("common/config.json")
	defer file.Close()

	if err != nil {
		log.Fatal("[loadConfig]: %s\n",err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)

	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n",err)
	}

	loadConfigFromEnvironment(&AppConfig)
}

func loadConfigFromEnvironment(AppConfig *configuration) {
	server, ok := os.LookupEnv("SERVER")

	if ok {
		AppConfig.Server = server
		log.Printf("[INFO]: Server information loaded from env.")
	}

	mysqlHost, ok := os.LookupEnv("MYSQLDB_HOST")

	if ok {
		AppConfig.MysqlDBHost = mysqlHost
		log.Printf("[INFO]: MYSQLDB server information loaded from env.")
	}



	mysqldbUser, ok := os.LookupEnv("MYSQLDB_USER")

	if ok {
		AppConfig.DBUser = mysqldbUser
		log.Printf("[INFO]: MYSQLDB USER information loaded from env.")
	}

	mysqldbPwd, ok := os.LookupEnv("MYSQLDB_PWD")

	if ok {
		AppConfig.DBPwd = mysqldbPwd
		log.Printf("[INFO]: MYSQLDB password information loaded from env.")
	}


	database, ok := os.LookupEnv("MYSQLDB_DATABASE")

	if ok {
		AppConfig.Database = database
		log.Printf("[INFO]: MYSQLDB database information loaded from env.")
	}

}

var db *sql.DB

func GetDBConnection() *sql.DB {
	if db == nil {
		var err error

		db, err = sql.Open("mysql","addnewuser:abc@tcp(198.168.1.1:3306)/mydb")

		if err != nil {
			log.Fatalf("[GetDBConnection]: %s\n", err)
		}

		if err := migrateDatabase(db); err!=nil {
			log.Fatalf("[MigrateDatabase]: %s\n", err)
		}

	}
	return db
}

func createDBConnection() {
	var err error

	db, err = sql.Open("mysql","addnewuser:abc@tcp(198.168.1.1:3306)/mydb")

	if err != nil {
		log.Fatalf("[CreateDBConnection]: %s\n", err)
	}

}

func migrateDatabase(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		return err
	}

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/common/migrations", dir),
		"mysql",
		driver,
	)

	if err != nil {
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("Applying database migrations")

	err = migration.Up()
	if err!=nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := migration.Version()
	if err != nil {
		return err
	}

	migration.Log.Printf("Active database version: %d", version)
	return nil

}
