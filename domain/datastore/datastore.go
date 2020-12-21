package datastore

import (
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
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
		time.Second,
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
