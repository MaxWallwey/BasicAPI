package db

import (
	"github.com/gocql/gocql"
	"log"
	"os"
	"time"
)

var (
	cassandraHost     = os.Getenv("CASSANDRA_HOST")
	cassandraUser     = os.Getenv("CASSANDRA_USER")
	cassandraPassword = os.Getenv("CASSANDRA_PASSWORD")
)

func SetupCassandra() *gocql.Session {
	cluster := gocql.NewCluster(cassandraHost)
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: 3,
	}
	cluster.Keyspace = "roster"
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraUser,
		Password: cassandraPassword,
	}
	var session *gocql.Session
	var err error
	for {
		session, err = cluster.CreateSession()
		if err == nil {
			break
		}
		log.Printf("CreateSession: %v", err)
		time.Sleep(time.Second)
	}
	log.Printf("Connected OK\n")
	return session
}
