package helpers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mayank:mayank@cluster0.dxf2y.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchList"

var collection *mongo.Collection

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established", client)

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection ready")
}
