package internal

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
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

func GetUser(login string) (string, int, error) {
	var hash string
	var id, num int
	row, err := dbPool.Query(context.Background(), "SELECT pass,id FROM users WHERE login=$1", login)
	for row.Next() {
		row.Scan(&hash, &id)
		num++
	}
	log.WithFields(log.Fields{
		"id":  id,
		"err": err,
	}).Debug("Auth_GetUser")
	if num == 0 {
		return "", 0, fmt.Errorf("no user with login: %v", login)
	}
	return hash, id, err
}
func AddUser(login, pass string) error {
	_, err := dbPool.Query(context.Background(), "INSERT INTO users (login,pass)VALUES ($1,$2)", login, pass)
	log.WithFields(log.Fields{
		"login": login,
		"err":   err,
	}).Debug("Auth_AddUser")
	return err
}
func UpdatePass(hash string, id int) error {
	_, err := dbPool.Query(context.Background(), "UPDATE users SET pass=$1 WHERE id=$2", hash, id)
	log.WithFields(log.Fields{
		"id":  id,
		"err": err,
	}).Debug("Auth_AddUser")
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
		user = "authuser"
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
		dbName = "authdb"
	}
	conURL = "postgres://" + user + ":" + pass + "@" + dbAdr + ":" + dbPort + "/" + dbName
}
