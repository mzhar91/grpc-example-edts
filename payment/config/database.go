package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

func InitDB() (dbConn *sql.DB) {
	var dsn string

	dbProvider := Env.DBProvider
	dbHost := Env.DBHost
	dbPort := Env.DBPort
	dbUser := Env.DBUser
	dbPass := Env.DBPasswd
	dbName := Env.DBName

	if dbProvider == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	} else if dbProvider == "mysql" {
		connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		val := url.Values{}
		val.Add("parseTime", "1")
		val.Add("charset", "utf8")
		dsn = fmt.Sprintf("%s?%s", connection, val.Encode())
	}

	dbConn, err := sql.Open(dbProvider, dsn)
	if err != nil && Env.Debug {
		fmt.Println(err)
	} else {
		dbConn.SetMaxIdleConns(10)
		dbConn.SetMaxOpenConns(100)
		dbConn.SetConnMaxLifetime(time.Minute * 4)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// defer dbConn.Close()

	return
}
