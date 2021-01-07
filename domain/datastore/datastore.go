package datastore

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"github.com/go-redis/redis"
	"time"
)

func NewCouchbaseSession() (*gocb.Cluster, error) {
	cluster, err := gocb.Connect(
		"localhost:8091",
		gocb.ClusterOptions{
			Username: "admin",
			Password: "adminadmin",
		},
	)
	if err != nil {
		panic(err)
	}
	err = cluster.WaitUntilReady(
		10*time.Second,
		&gocb.WaitUntilReadyOptions{
			DesiredState: gocb.ClusterStateOnline,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("connecting to couchbase databse ...")
	_, err = cluster.Ping(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("db ping success ")

	return cluster, nil
}

func NewRedisDbConnection() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redisDB

}
