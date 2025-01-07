package connecter

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectWithMongo(databaseName, collectionName string) *mongo.Collection {
	env := os.Getenv("GO_ENV")

	var envFile string
	if env == "development" {
		envFile = ".env"
	} else {
		envFile = ".env.prod"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %v file : %v", envFile, err)
	}

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")

	conn := "mongodb+srv://" + mongoUsername + ":" + mongoPassword + "@user-cluster.xgu8qno.mongodb.net/?retryWrites=true&w=majority&appName=user-cluster"

	client, err := mongo.Connect(options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	collection := client.Database(databaseName).Collection(collectionName)
	log.Printf("Running in %s environment", env)

	return collection
}
