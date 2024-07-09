package helper
import (
	"context"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"github.com/cerdas-buatan/be/model"
)

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

	return client.Database("chatbot")
}

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

func InsertTwoDoc(database *mongo.Database, collection string, document interface{}) (InsertedID interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := database.Collection(collection).InsertOne(ctx, document)
	if err != nil {
		return nil
	}

	return result.InsertedID
}

// InsertUserdata inserts a new user into the database
func InsertUserdata(database *mongo.Database, username, email, password, salt, role string) (InsertedID interface{}) {
	// Create the User struct
	user := model.User{
		Email:    email,
		Password: password,
		Salt:     salt,
		Role:     role,
	}

	// Create the Pengguna struct
	pengguna := model.Pengguna{
		Username: username,
		Akun:     user,
	}

	return InsertTwoDoc(database, "users", pengguna)
}


func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (interface{}, error) {
	coll := db.Collection(collection)
	result, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}


// func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
// 	result, err := db.Collection(col).InsertOne(context.Background(), doc)
// 	if err != nil {
// 		return insertedID, fmt.Errorf("kesalahan server : insert")
// 	}
// 	insertedID = result.InsertedID.(primitive.ObjectID)
// 	return insertedID, nil
// }

func FindOneDoc(db *mongo.Database, collection string, filter bson.M) *mongo.SingleResult {
	return db.Collection(collection).FindOne(nil, filter)
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}