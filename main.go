package main

import (
	"log"
	"medikakh/controller"
	"medikakh/domain/datastore"
)

func main() {
	rdConnection := datastore.NewRedisDbConnection()
	err := rdConnection.Ping().Err()
	if err != nil {
		log.Println("error on ping with redis : " + err.Error())
	}

	controller.Run(":50501")
}
