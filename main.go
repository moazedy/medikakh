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

	cluster, err := datastore.NewCouchbaseSession()
	if err != nil {
		log.Fatal("error on creating couchbase session")
	}
	_, err = cluster.Ping(nil)
	if err != nil {
		panic(err)
	}
	log.Println("db ping success ")

	controller.Run(":50501")
}
