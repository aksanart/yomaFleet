package mongodb

import (
	"context"
	"time"

	"github.com/aksanart/vehicle/model"
	"github.com/aksanart/vehicle/pkg/constant"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection = "vehicle"

type MongoClient struct {
	conn *mongo.Client
	db   string
}

func (r *MongoClient) HealthCheck(ctx context.Context) error {
	return r.conn.Ping(ctx, nil)
}

func (r *MongoClient) FindAllVehilce(ctx context.Context, page int) (res []*model.Vehicle, err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{}
	// Define the offset and limit
	offset := int64(page)
	limit := int64(constant.MONGO_LIMIT_DISPLAY)

	// Fetch data with offset and limit
	findOptions := options.Find()
	findOptions.SetSkip(offset)
	findOptions.SetLimit(limit)

	data, err := coll.Find(ctx, filter, findOptions)
	if err != nil {
		if err.Error() == constant.MONGO_NO_DATA_ALL || err.Error() == constant.MONGO_NO_DATA_SINGLE {
			return res, nil
		}
		return
	}
	defer data.Close(ctx)

	for data.Next(ctx) {
		var rowCat *model.Vehicle
		err = data.Decode(&rowCat)
		if err != nil {
			return
		}
		res = append(res, rowCat)
	}
	return
}

func (r *MongoClient) FindById(ctx context.Context, id string) (res *model.Vehicle, err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"id": id}
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

func (r *MongoClient) Create(ctx context.Context, data *model.Vehicle) (id string, err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	id = uuid.NewString()
	data.ID = id
	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = time.Now().Unix()
	bsonMap, err := prepareData(data)
	if err != nil {
		return
	}
	_, err = coll.InsertOne(ctx, bsonMap)
	if err != nil {
		return
	}
	return 
}

func (r *MongoClient) Update(ctx context.Context, id string, data *model.Vehicle) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"id": id}
	data.UpdatedAt = time.Now().Unix()
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

func (r *MongoClient) Delete(ctx context.Context, id string) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"id": id}
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
