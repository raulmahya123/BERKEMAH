package belajarmahya

import (
	"context"
	"os"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetConnection(mongoenvkatalogfilm, dbname string) *mongo.Database {
	var DBmongoinfo = atdb.DBInfo{
		DBString: os.Getenv(mongoenvkatalogfilm),
		DBName:   dbname,
	}
	return atdb.MongoConnect(DBmongoinfo)
}

//------------------------------------------------------------------- User

// Create

func InsertUser(mconn *mongo.Database, collname string, datauser User) interface{} {
	return atdb.InsertOneDoc(mconn, collname, datauser)
}

// Read

func GetAllUser(mconn *mongo.Database, collname string) []User {
	user := atdb.GetAllDoc[[]User](mconn, collname)
	return user
}

func FindUser(mconn *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"username": userdata.Username}
	return atdb.GetOneDoc[User](mconn, collname, filter)
}

func FindPassword(mconn *mongo.Database, collname string, userdata User) User {
	filter := bson.M{"password": userdata.Password}
	return atdb.GetOneDoc[User](mconn, collname, filter)
}

func IsPasswordValid(mconn *mongo.Database, collname string, userdata User) bool {
	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, collname, filter)
	hashChecker := CheckPasswordHash(userdata.Password, res.Password)
	return hashChecker
}

func UsernameExists(mongoenvkatalogfilm, dbname string, userdata User) bool {
	mconn := SetConnection(mongoenvkatalogfilm, dbname).Collection("user")
	filter := bson.M{"username": userdata.Username}

	var user User
	err := mconn.FindOne(context.Background(), filter).Decode(&user)
	return err == nil
}

// Update

func EditUser(mconn *mongo.Database, collname string, datauser User) interface{} {
	filter := bson.M{"username": datauser.Username}
	return atdb.ReplaceOneDoc(mconn, collname, filter, datauser)
}

// Delete

func DeleteUser(mconn *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"username": userdata.Username}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}

//forgot password

func ForgotPassword(mconn *mongo.Database, collname string, userdata User) interface{} {
	filter := bson.M{"password": userdata.Password}
	return atdb.ReplaceOneDoc(mconn, collname, filter, userdata)
}

//------------------------------------------------------------------- ProductOnlineCourse

func InsertProductOnlineCourse(mconn *mongo.Database, collname string, data ProductOnlineCourse) interface{} {
	return atdb.InsertOneDoc(mconn, collname, data)
}

func GetAllProductOnlineCourse(mconn *mongo.Database, collname string) []ProductOnlineCourse {
	data := atdb.GetAllDoc[[]ProductOnlineCourse](mconn, collname)
	return data
}

func FindProductOnlineCourse(mconn *mongo.Database, collname string, data ProductOnlineCourse) ProductOnlineCourse {
	filter := bson.M{"id": data.ID}
	return atdb.GetOneDoc[ProductOnlineCourse](mconn, collname, filter)
}

func EditProductOnlineCourse(mconn *mongo.Database, collname string, data ProductOnlineCourse) interface{} {
	filter := bson.M{"id": data.ID}
	return atdb.ReplaceOneDoc(mconn, collname, filter, data)
}

func DeleteProductOnlineCourse(mconn *mongo.Database, collname string, data ProductOnlineCourse) interface{} {
	filter := bson.M{"id": data.ID}
	return atdb.DeleteOneDoc(mconn, collname, filter)
}
