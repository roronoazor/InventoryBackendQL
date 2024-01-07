package database

import (
	"context"
	"log"

	config "github.com/roronoazor/InventoryBackendQL/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var MongoDbInstance MongoInstance

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/docs/drivers/go/current/quick-start/
func Connect() error {

	dbName := config.Config("DB_NAME")
	mongoURI := config.Config("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database(dbName)

	if err != nil {
		return err
	}

	MongoDbInstance = MongoInstance{
		Client: client,
		Db:     db,
	}

	log.Println("Connected Successfully")
	return nil
}
