package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var mongoDBURL string
	var isAuth bool
	if username == "" && password == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	// Create a new client and connect to the server
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failde to connect to mongoDB due to error: %v", err)
	}

	// Send a ping to confirm a successful connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failde to ping mongoDB due to error: %v", err)
	}

	return client.Database(database), nil
}
