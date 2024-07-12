package config

import (
	"os"

	"github.com/gocroot/helper/atdb"
)

// connection mongo
var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "gaysdisal_db",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)