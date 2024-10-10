package mongodb

import (
	"context"
	"time"

	"github.com/aksan/weplus/apigw/model"
	"github.com/aksan/weplus/apigw/pkg/constant"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "user"

type MongoClient struct {
	conn *mongo.Client
	db   string
}

func (r *MongoClient) HealthCheck(ctx context.Context) error {
	return r.conn.Ping(ctx, nil)
}

func (r *MongoClient) FindByEmail(ctx context.Context, email string) (res *model.User, err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"email": email}
	data := coll.FindOne(ctx, filter)
	err = data.Decode(&res)
	if err != nil {
		if err.Error() == constant.MONGO_NO_DATA_ALL || err.Error() == constant.MONGO_NO_DATA_SINGLE {
			return res, nil
		}
		return
	}
	return
}

func (r *MongoClient) Create(ctx context.Context, data *model.User) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	data.ID = uuid.NewString()
	data.CreatedAt = time.Now().Unix()
	bsonMap, err := prepareData(data)
	if err != nil {
		return
	}
	_, err = coll.InsertOne(ctx, bsonMap)
	if err != nil {
		return
	}
	return nil
}

func (r *MongoClient) Update(ctx context.Context, email string, data *model.User) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"email": email}
	bsonMap, err := prepareData(data)
	if err != nil {
		return
	}
	update := bson.M{"$set": bsonMap}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}
	return nil
}

func (r *MongoClient) Delete(ctx context.Context, email string) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"email": email}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return
	}
	return nil
}

func prepareData(data interface{}) (bsonMap bson.M, err error) {
	bsonData, err := bson.Marshal(data)
	if err != nil {
		return
	}
	err = bson.Unmarshal(bsonData, &bsonMap)
	if err != nil {
		return
	}
	return
}
