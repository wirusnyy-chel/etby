package internal

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v4/pgxpool"
)

// var conURL = "postgres://authuser:auth@localhost:5432/authdb"
var conURL string
var dbPool *pgxpool.Pool

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	initURI()
	log.WithField("URI", conURL).Info("Connecting to db")
	var err error
	poolConf, _ := pgxpool.ParseConfig(conURL)
	dbPool, err = pgxpool.ConnectConfig(context.Background(), poolConf)
	if err != nil {
		log.Panic(err)
	}
}

func GetUser(login string) (string, int, error) {
	var hash string
	var id int
	err := dbPool.QueryRow(context.Background(), "SELECT pass,id FROM users WHERE login=$1", login).Scan(&hash, &id)
	log.WithFields(log.Fields{
		"id":  id,
		"err": err,
	}).Debug("Auth_GetUser")
	return hash, id, err
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
