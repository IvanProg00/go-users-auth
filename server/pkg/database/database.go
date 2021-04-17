package database

import (
	"context"
	"fmt"
	"os"
	"time"
	"users-authentication/pkg/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseUsers = "users_auth"

	CollectionUsers = "users"
)

var MI MongoInstance

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewConnection() error {
	username := os.Getenv(util.MongoUsername)
	if username == "" {
		username = "mongo"
	}
	password := os.Getenv(util.MongoPassword)
	if password == "" {
		password = "pass"
	}
	host := os.Getenv(util.MongoHost)
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv(util.MongoPort)
	if port == "" {
		port = "27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port))
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		return err
	}
	MI.Client = client
	return nil
}

func NewConnectionDatabase(database string) error {
	err := NewConnection()
	if err != nil {
		return err
	}

	MI.DB = MI.Client.Database(database)
	return nil
}