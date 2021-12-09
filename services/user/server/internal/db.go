package internal

import (
	"context"
	"os"

	userPB "etby/services/user"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

//var conURL = "postgres://authuser:auth@localhost:5432/authdb"
var conURL string
var dbPool *pgxpool.Pool

func init() {
	initURI()
	log.WithField("URI", conURL).Info("Connecting to db")
	var err error
	dbPool, err = pgxpool.Connect(context.Background(), conURL)
	if err != nil {
		log.Panic(err)
	}
}

func GetUser(id int64) (int64, string, userPB.Privacy, error) {
	var name string
	var privacy userPB.Privacy
	row, err := dbPool.Query(context.Background(), "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return 0, "", 0, err
	}
	for row.Next() {
		err = row.Scan(&id, &name, &privacy)
	}
	log.WithFields(log.Fields{
		"id":      id,
		"name":    name,
		"privacy": privacy,
		"err":     err,
	}).Debug("User_GetUser")
	return id, name, privacy, err
}
func AddUser(id int64, name string, privacy userPB.Privacy) error {
	_, err := dbPool.Exec(context.Background(), "INSERT INTO users (id,name,privacy)VALUES ($1,$2,$3)", id, name, privacy)
	log.WithFields(log.Fields{
		"id":      id,
		"name":    name,
		"privacy": privacy,
		"err":     err,
	}).Debug("User_GetUser")
	return err
}
func UpdateUser(id int64, name string, privacy userPB.Privacy) error {
	_, err := dbPool.Exec(context.Background(), "UPDATE users SET (name=$1,privacy=$2) WHERE id=$3", name, privacy, id)
	log.WithFields(log.Fields{
		"id":      id,
		"name":    name,
		"privacy": privacy,
		"err":     err,
	}).Debug("User_GetUser")
	return err
}
func CheckConn() {
	err := dbPool.Ping(context.Background())
	if err != nil {
		log.Panic(err)
	}
	log.Info("Successfully connected to db")
}
func initURI() {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "user"
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "auth"
	}
	dbAdr := os.Getenv("DB_ADR")
	if dbAdr == "" {
		dbAdr = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "userdb"
	}
	conURL = "postgres://" + user + ":" + pass + "@" + dbAdr + ":" + dbPort + "/" + dbName
}
