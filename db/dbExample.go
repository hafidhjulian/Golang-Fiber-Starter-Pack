package db

import (
	"fmt"
	"golang-fiber-starterpack/config"
	"log"

	"xorm.io/xorm"
)

func ConnectDB() *xorm.Engine {
	config := config.LoadConfig(".")

	engine, err := xorm.NewEngine("postgres", config.DbDsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgresql success")
	return engine
}
