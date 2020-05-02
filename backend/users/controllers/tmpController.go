package controllers

import (
	"html/template"
	"os"
	"log"
	"net/http"

	"github.com/shubham1010/project/backend/users/models"
	"github.com/shubham1010/project/backend/users/common"

	_ "github.com/go-sql-driver/mysql"
//	"github.com/dgrijalva/jwt-go"
)

//var jwtKey = []byte("my_secret_key")

/*type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
*/
var tpl *template.Template

func Init() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

//	tpl, _ = template.ParseFiles(dir+"/controllers/views/index.html")
	tpl, err = template.ParseGlob(dir+"/containerData/views/*.html")

	if err != nil {
		log.Fatal("Error loading templates: "+ err.Error())
		 //w.WriteHeader(http.StatusInternalServerError)
	}
}

func  Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w,"login.html", nil)
}

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	newUser := &models.UserSignUp{}
	newUser.Username = r.FormValue("username")
	newUser.Password = r.FormValue("password")

	if newUser.Username=="" || newUser.Password=="" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	db := common.GetDBConnection()

	row := db.QueryRow("SELECT * FROM userdata where username=?",newUser.Username)

	dbData := models.UserSignUp{}

	err := row.Scan(&dbData.Username, &dbData.Password)

	if err != nil {
		common.DisplayAppError(w, err, "Row Scanning Failed",500)
		http.Redirect(w,r,"/login",http.StatusSeeOther)
	}


	if dbData.Username == newUser.Username && dbData.Password == newUser.Password {
		tpl.ExecuteTemplate(w,"user.html", newUser)
		/*expirationTime := time.Now().Add(5 * time.Minutes)
		claims := &Claims {
			Username: dbData.Username,
			StandardClaims: jwt.StandardClaims {
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Cookie(w, &http.Cookie{
			Name:		"token",
			Value:		tokenString,
			Expires:	expirationTime,
		})
		*/
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func NewUser(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w,"signUp.html",nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	newUser := &models.UserSignUp{}
	newUser.Username = r.FormValue("username")
	newUser.Password = r.FormValue("password")


	if newUser.Username =="" || newUser.Password == "" {
		http.Redirect(w, r, "/newuser", http.StatusSeeOther)
		return
	}

	dbConn := common.GetDBConnection()

	stmt, err := dbConn.Prepare("INSERT INTO userdata VALUES(?,?)")

	if err!= nil {
		common.DisplayAppError(w, err, "Database Insertion Failed",500)
		return
	}
	_, err = stmt.Exec(newUser.Username,newUser.Password)
	if err != nil {
		panic(err.Error())
		log.Fatal("Exec Statement fails")
	}
	tpl.ExecuteTemplate(w,"gotNewUser.html",newUser)
}
