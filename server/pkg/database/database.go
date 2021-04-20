package database

import (
	"context"
	"fmt"
	"os"
	"users-authentication/pkg/configs"
	"users-authentication/pkg/util"

	"go.mongodb.org/mongo-driver/bson"
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

	ctx, cancel := context.WithTimeout(context.Background(), configs.TimeToConnectDB)
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

	if _, err := ConfigDatabase(); err != nil {
		return err
	}
	return nil
}

func ConfigDatabase() (string, error) {
	return MI.DB.
		Collection(CollectionUsers).
		Indexes().
		CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bson.D{{Key: "password", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
}
