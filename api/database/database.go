package database

import (
	"TheWarEconomy/api/utils"
	"fmt"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/joho/godotenv"
)

var usersCol *gocb.Collection
var pmcsCol *gocb.Collection
var tweScope *gocb.Scope

func init() {
	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Println("Error getting db vars from env", envErr)
		os.Exit(1)
	}

	// @TODO: Additional security. Keys?
	cluster, err := gocb.Connect(os.Getenv(utils.EnvDbUrl), gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: os.Getenv(utils.EnvDbUser),
			Password: os.Getenv(utils.EnvDbPassword),
		},
	})

	if err != nil {
		fmt.Println("Error connecting to cluster", err)
		os.Exit(1)
	}

	bucket := cluster.Bucket("the-war-economy")
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		fmt.Println("Error connecting to bucket", err)
		os.Exit(1)
	}

	scope := bucket.Scope("_default")

	uc := scope.Collection("users")
	pmcc := scope.Collection("pmcs")

	if uc != nil && pmcc != nil {
		usersCol = uc
		pmcsCol = pmcc
		tweScope = scope
	} else {
		fmt.Println("Error getting collections. Terminating...")
		os.Exit(1)
	}
}

func GetUsersCol() *gocb.Collection {
	return usersCol
}

func GetPmcsCol() *gocb.Collection {
	return pmcsCol
}

func GetTweScope() *gocb.Scope {
	return tweScope
}
