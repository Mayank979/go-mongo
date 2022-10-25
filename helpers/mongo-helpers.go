package helpers

import (
	"context"
	"fmt"
	"log"

	model "github.com/mayankyadav/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOne(movie model.Netflix) {
	result, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result is", result, result.InsertedID)

}

func UpdateOne(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("update one result is", result)

}

func DeleteOne(movieId string) int {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("delete result is", result)

	return int(result.DeletedCount)

}

func DeleteAll() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All records deleted result is", result)

	return result.DeletedCount
}

func GetAllItems() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}
