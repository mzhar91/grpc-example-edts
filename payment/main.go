package main

import (
	"fmt"
	"log"
	"os"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	
	_grpc "github.com/grpc-example-edts/payment/cmd"
	_config "github.com/grpc-example-edts/payment/config"
	_load "github.com/grpc-example-edts/payment/config/load"
)

func init() {
	log.SetFlags(log.Flags() | log.Llongfile)
	log.SetOutput(os.Stdout)
	
	_config.LoadEnv()
	if _config.Env.Debug {
		fmt.Println("Service RUN on DEBUG mode")
	}
	
}

func main() {
	dbConn := _config.InitDB()
	defer dbConn.Close()
	
	connection := _config.Connection{Database: dbConn}
	
	e := echo.New()
	
	// Get timeoutcontext
	timeoutContext := _config.GetTimeoutContext()
	
	_load.Load(e, &connection, timeoutContext)
	
	_config.ApiSetup()
	
	go func() {
		_grpc.StartServerGRPC(&connection, timeoutContext)
	}()
	
	e.Start(_config.Env.ServerAddr)
}
