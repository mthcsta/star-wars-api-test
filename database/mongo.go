package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mthcsta/star-wars-api-test/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(common.Config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MONGODB: connected with host")

	// check if database selected exists
	listDatabaseNames, err := client.ListDatabaseNames(ctx, bson.M{})
	databaseExists := false
	for _, databaseName := range listDatabaseNames {
		if databaseName == common.Config.DbName {
			databaseExists = true
			break
		}
	}
	if databaseExists == false {
		log.Fatalf("MONGODB: database '%s' not found", common.Config.DbName)
		os.Exit(-1)
	}

	fmt.Printf("MONGODB: connected to database '%s'\n", common.Config.DbName)

	// check if database selected has all collections necessary
	clientDatabase := client.Database(common.Config.DbName)
	listCollectionNames, err := clientDatabase.ListCollectionNames(ctx, bson.M{})
	var collectionExists bool
	for _, collection := range ListCollections {
		collectionExists = false
		for _, collectionName := range listCollectionNames {
			if collection == collectionName {
				collectionExists = true
				break
			}
		}
		if collectionExists == false {
			fmt.Printf("MONGODB: creating collection '%s' in database '%s'\n", collection, common.Config.DbName)
			clientDatabase.CreateCollection(ctx, collection)
		}
	}

	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()
var ListCollections []string = []string{"planets", "films"}

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(common.Config.DbName).Collection(collectionName)
	return collection
}
