package database

import (
	"fmt"
	"log"
	"time"
	"os"
	"context"
	// "path/filepath"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{

	err := godotenv.Load(".env")
	// err := godotenv.Load(filepath.Join(base, ".env"))
	// err := godotenv.Load(filepath.Abs(".env"))
	// path := filepath.Join("github.com/lordscoba/golang-auth/pkg/helpers","file.txt")
	// fmt.Println(path)
	if err != nil {
		log.Println("Errpor loaing .env file")
	}


	MongoDb := os.Getenv("MONGODB_URL")

	client , err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb!")

	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("scribbles").Collection(collectionName)
	return collection
}
