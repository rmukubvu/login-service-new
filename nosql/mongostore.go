package nosql

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type AuthStore struct {
	client *mongo.Client
	db     *mongo.Database
}

const (
	schemaName = "amakosi_users"
)

func NewConnection(uri string) *AuthStore {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &AuthStore{db: client.Database(schemaName),
		client: client}
}

func (as *AuthStore) InsertRecord(documentName string, data interface{}) error {
	collection := as.db.Collection(documentName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//then insert
	_, err := collection.InsertOne(ctx, data)
	return err
}

func (as *AuthStore) SingleRecord(documentName string, filter bson.D) (*mongo.SingleResult, error) {
	collection := as.db.Collection(documentName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	singleResult := collection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}
	return singleResult, nil
}

func (as *AuthStore) CloseDb() {
	if as.client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		as.client.Disconnect(ctx)
	}
}
