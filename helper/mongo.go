package helper

import (
	"context"
	"fmt"
	"os"
	"time"
	"regexp"
	"strings"
	"github.com/RadhiFadlillah/go-sastrawi"

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

func QueriesDataRegexp(db *mongo.Database, ctx context.Context, queries string) (dest model.Dataset, err error) {
	filter := bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
	err = db.Collection("dataset").FindOne(ctx, filter).Decode(&dest)

	if err != nil && err != mongo.ErrNoDocuments {
		return dest, err
	}

	return dest, err
}

func QueriesSecret(db *mongo.Database, ctx context.Context, secret string) (dest model.Secrets, err error) {
	filter := bson.M{"secret_token": primitive.Regex{Pattern: secret, Options: "i"}}
	err = db.Collection("Secret").FindOne(ctx, filter).Decode(&dest)

	if err != nil && err != mongo.ErrNoDocuments {
		return dest, err
	}

	return dest, err
}

func QueriesDataRegexpALL(db *mongo.Database, ctx context.Context, queries string) (dest model.Dataset, score float64, err error) {
	queries = SeparateSuffixMu(queries)
	var cursor *mongo.Cursor
	queries = Stemmer(queries)
	splits := strings.Split(queries, " ")

	if len(splits) >= 5 {
		queries = splits[len(splits)-3] + " " + splits[len(splits)-2] + " " + splits[len(splits)-1]
		filter := bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
		cursor, err = db.Collection("dataset").Find(ctx, filter)

		if err != nil && err != mongo.ErrNoDocuments {
			queries = splits[len(splits)-2] + " " + splits[len(splits)-1]
			filter = bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
			cursor, err = db.Collection("dataset").Find(ctx, filter)
			if err != nil && err != mongo.ErrNoDocuments {
				queries = splits[len(splits)-1]
				filter = bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
				cursor, err = db.Collection("dataset").Find(ctx, filter)
				if err != nil && err != mongo.ErrNoDocuments {
					filter = bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
					cursor, err = db.Collection("dataset").Find(ctx, filter)
					if err != nil && err != mongo.ErrNoDocuments {
						return dest, score, err
					}
				}
			}
		}
	} else if len(splits) == 1 {
		queries = splits[0]
		filter := bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
		cursor, err = db.Collection("dataset").Find(ctx, filter)
	} else if len(splits) <= 4 {
		queries = splits[len(splits)-2] + " " + splits[len(splits)-1]
		filter := bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
		cursor, err = db.Collection("dataset").Find(ctx, filter)

		if err != nil && err != mongo.ErrNoDocuments {
			queries = splits[len(splits)-1]
			filter = bson.M{"questions": primitive.Regex{Pattern: queries, Options: "i"}}
			cursor, err = db.Collection("dataset").Find(ctx, filter)
			if err != nil && err != mongo.ErrNoDocuments {
				return dest, score, err
			}
		}
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var data model.Dataset
		err := cursor.Decode(&data)
		if err != nil {
			return data, score, err
		}
		// Calculate BoW similarity score
		str2 := data.Question
		scorex := BagOfWordsSimilarity(queries, str2)
		if score < scorex {
			dest = data
			score = scorex
		}
	}

	return dest, score, err
}

func QueriesALL(db *mongo.Database, ctx context.Context) (dest []model.Dataset, err error) {
	filter := bson.M{}
	cursor, err := db.Collection("dataset").Find(ctx, filter)

	if err != nil && err != mongo.ErrNoDocuments {
		return dest, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var data model.Dataset
		err := cursor.Decode(&data)
		if err != nil {
			return nil, err
		}
		dest = append(dest, data)
	}

	return dest, err
}

func Stemmer(Sentences string) (newString string) {
	dictionary := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dictionary)
	for _, word := range sastrawi.Tokenize(Sentences) {
		//fmt.Println(word)
		newString = newString + " " + stemmer.Stem(word)
		//fmt.Println(newString)
	}
	return strings.TrimSpace(newString)
}

func SeparateSuffixMu(word string) string {
	// Regex for detecting words with suffixes "mu" and "ku" at the end
	re := regexp.MustCompile(`(\w+)(mu|ku)$`)

	// Check if the word matches the regex
	if re.MatchString(word) {
		// Replace "mu" and "ku" with "kamu" and "aku" respectively
		return re.ReplaceAllStringFunc(word, func(matched string) string {
			if matched[len(matched)-2:] == "mu" {
				return matched[:len(matched)-2] + " kamu"
			}
			return matched[:len(matched)-2] + " aku"
		})
	}

	// If there are no suffixes "mu" or "ku", return the original word
	return word
}