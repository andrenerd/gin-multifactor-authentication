package main

import (
	_ "fmt"
	"reflect"
	"encoding/json"
	"io/ioutil"
        _ "github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/doug-martin/goqu/v9"
	"github.com/andrenerd/gin-multifactor-authentication"
)

var schema = `
CREATE TABLE IF NOT EXISTS user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username STRING(64),
        email STRING,
        password STRING(64)
);

CREATE TABLE IF NOT EXISTS user_service_authenticator (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL UNIQUE,
        key STRING,
        FOREIGN KEY (user_id) REFERENCES user(id)
);
`

var db, err = sqlx.Connect("sqlite3", ":memory:")

type User struct {
	multauth.User
	db *sqlx.DB

	Id       int64  `db:"id" json:"id" goqu:"skipinsert"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

func (user *User) GetByIdentifier(identifier string, value interface{}) error {
	t := reflect.TypeOf(*user)
	f, _ := t.FieldByName(identifier)

	dbField, ok := f.Tag.Lookup("db")
	if !ok {
		dbField = identifier
	}

	query, _, _ := goqu.
		From("user").
		Where(goqu.Ex{dbField: value}).
		ToSQL()

	err := user.db.Get(user, query)
	return err
}

func (user *User) GetServices() ([]multauth.ServiceInterface, error) {
	services := []multauth.ServiceInterface{&AuthenticatorService{}}

	query, _, _ := goqu.
		From("user_service_authenticator").
		Where(goqu.Ex{"Id": user.Id}).
		ToSQL()

	err := user.db.Get(services[0], query)
	return services, err
}

func (user *User) Save(fields ...[]string) error {
	var query string

	if user.Id == 0 {
		query, _, _ = goqu.
			Insert("user").
			Cols("username").
			Rows(user).
			ToSQL()
	} else {
		query, _, _ = goqu.
			Update("user").
			Set(user).
			Where(goqu.Ex{"Id": user.Id}).
			ToSQL()
	}

	result, err := user.db.Exec(query)
	if err == nil {
		if id, err := result.LastInsertId(); err == nil {
			user.Id = id
		}
	}

	return err
}

type AuthenticatorService struct {
	multauth.AuthenticatorService
	db *sqlx.DB

	Id     int64 `db:"id" json:"id" goqu:"skipinsert"`
	UserId int64 `db:"user_id" json:"id"`
}

func (service *AuthenticatorService) Save(fields ...[]string) error {
	var query string

	if service.Id == 0 {
		query, _, _ = goqu.
			Insert("user_service_authenticator").
			Rows(service).
			ToSQL()
	} else {
		query, _, _ = goqu.
			Update("user_service_authenticator").
			Set(service).
			Where(goqu.Ex{"Id": service.Id}).
			ToSQL()
	}

	result, err := service.db.Exec(query)
	if err == nil {
		if id, err := result.LastInsertId(); err == nil {
			service.Id = id
		}
	}

	return err
}

func init() {
	db.MustExec(schema)

	// Seed user
	user := &User{db: db, Username: "johndoe", Email: "johndoe@email.com"}
	user.SetPassword("password")
	user.Save()

	// Seed user service
	service := &AuthenticatorService{db: db, UserId: user.Id}
	service.Init(map[string]interface{}{"Issuer": "Multauth", "AccountName": user.Username})
	service.Save()
}

func main() {
	auth := multauth.Auth{
		Flows: []multauth.Flow{{"Username", "Password", "Passcode"}},
	}

	app := gin.Default()

	app.POST("/signin", func(c *gin.Context) {
		var user = &User{db: db}
		var data map[string]interface{}

		byteData, _ := ioutil.ReadAll(c.Request.Body)
		if err := json.Unmarshal(byteData, &data); err != nil {
			c.Next()
		}

		err := auth.Authenticate(map[string]interface{}{
			"Username": data["username"],
			"Email":    data["email"],
			"Password": data["password"],
			"Passcode": data["passcode"],
		}, user)

		if err == nil {
			c.JSON(200, gin.H{
				"message": "Welcome " + user.Username,
				"token":   "YOUR_JWT_TOKEN",
			})
		} else {
			c.JSON(401, gin.H{
				"message": "I don't know you",
				"error":   err,
			})
		}
	})

	app.Run()
}
