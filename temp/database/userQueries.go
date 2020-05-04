package database

import (
	"database/sql"

	"github.com/shubham1010/project/temp/models"
	"github.com/shubham1010/project/temp/password"

	log "github.com/sirupsen/logrus"
)

func CheckUserExist(db *sql.DB, newUser models.User) (models.User) {
	log.Info("[DATABASE]: Inside CheckUserExist")
	dbUser := models.User{}
	row := db.QueryRow("SELECT username, password FROM users where username=?",newUser.Username)

    err := row.Scan(&dbUser.Username,&dbUser.Pass)
	if err!=nil {
		dbUser.Username=" "
		dbUser.Pass=" "
	}

	return dbUser
}

func InsertNewUser(db *sql.DB, newUser models.User) (error) {
	log.Info("[DATABASE]: Inside InsertNewUser")

	PassHash, err := password.HashPassword(newUser.Pass) 
	if err!=nil {
		log.Fatal("[DATABASE]: InsertNewUser Password is not Hashed")
	}

	newUser.Pass = PassHash

	stmt, err := db.Prepare("INSERT INTO users(id, username, password, name, email) VALUES(?,?,?,?,?)")
    if err != nil {
        log.Fatal("[DATABASE]: InsertNewUser Failed ", err)
    }

    _, err = stmt.Exec(newUser.Id, newUser.Username, newUser.Pass, newUser.Name, newUser.Email)

	return err
}
