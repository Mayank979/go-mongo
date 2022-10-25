package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mayankyadav/mongodb/router"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {
	fmt.Println("Database with Go")
	fmt.Println("Server is starting...")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))

	fmt.Println("Listening at port 3000")

}
