package apiserver

import (
	"context"
	"github.com/turopvin/go-rest-api/internal/app/store/mongostore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Start(config *Config) error {
	client, err := createDbClient(config)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	store := mongostore.New(client)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func createDbClient(config *Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DatabaseURL))
	if err != nil {
		return nil, err
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return client, nil
}
