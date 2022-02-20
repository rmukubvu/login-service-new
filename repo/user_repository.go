package repo

import (
	"go.mongodb.org/mongo-driver/bson"
	"login-service/config"
	"login-service/model"
	"login-service/nosql"
)

var ds *nosql.AuthStore

func init() {
	ds = nosql.NewConnection(config.MongoUrlConfig().Url)
}

func Auth(user model.User) error {
	err := ds.InsertRecord("user", user)
	if err != nil {
		return err
	}
	return nil
}

func LoginWith(username string) (model.User, error) {
	filter := bson.D{{"user_name", username}}
	result, err := ds.SingleRecord("user", filter)
	if err != nil {
		return model.User{}, err
	}
	user := model.User{}
	err = result.Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func CloseDb() {
	ds.CloseDb()
}
