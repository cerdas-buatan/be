package mongo
import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
s
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	model "github.com/cerdas-buatan/be/model"
	atdb "github.com/aiteung/atdb"
)

// connection db
func ConnectDB(uri string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbName)
}

// mongo connec
func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

// Insert a document into a collection
func InsertDoc(db *mongo.Database, col string, doc interface{}) (interface{}, error) {
	collection := db.Collection(col)
	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, fmt.Errorf("error InsertDoc %s: %s", col, err)
	}
	return result.InsertedID, nil
}

// Update a document in a collection
func UpdateDoc(db *mongo.Database, col string, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.Collection(col)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("error UpdateDoc %s: %s", col, err)
	}
	return result, nil
}

// Delete a document from a collection
func DeleteDoc(db *mongo.Database, col string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.Collection(col)
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, fmt.Errorf("error DeleteDoc %s: %s", col, err)
	}
	return result, nil
}