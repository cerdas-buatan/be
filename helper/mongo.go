package helper

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	model "github.com/cerdas-buatan/be/model"
	atdb "github.com/aiteung/atdb"
)

// mongo connect
func MongoConnect(MONGOSTRING, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRING)))
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

func SetConnection() *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv("MONGOSTRING"),
		DBName:   "AI",
	}
	return atdb.MongoConnect(DBmongoinfo)
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

func InsertUser(db *mongo.Database, collection string, userdata model.User) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "username : " + userdata.Email + " password : " + userdata.Password
}

// InsertUserdata insert user untuk register
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

func FindOneDoc(db *mongo.Database, collection string, filter bson.M) *mongo.SingleResult {
	return db.Collection(collection).FindOne(nil, filter)
}

// updateonedoc
func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) error {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("tidak ada data yang diubah")
	}
	return nil
}

func FindUserByUsername(db *mongo.Database, username string) (model.Pengguna, error) {
	pengguna := model.Pengguna{}
	filter := bson.M{"username": username}
	err := db.Collection("pengguna").FindOne(context.Background(), filter).Decode(&pengguna)
	if err != nil {
		return model.Pengguna{}, err
	}
	return pengguna, nil
}
