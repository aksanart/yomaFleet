package mongodb

import (
	"context"
	"time"

	"github.com/aksanart/tracker_service/model"
	"github.com/aksanart/tracker_service/pkg/constant"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "tracker"

type MongoClient struct {
	conn *mongo.Client
	db   string
}

func (r *MongoClient) HealthCheck(ctx context.Context) error {
	return r.conn.Ping(ctx, nil)
}

func (r *MongoClient) FindAllTracker(ctx context.Context) (res []*model.Tracker, err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{}
	data, err := coll.Find(ctx, filter)
	if err != nil {
		if err.Error() == constant.MONGO_NO_DATA_ALL || err.Error() == constant.MONGO_NO_DATA_SINGLE {
			return res, nil
		}
		return
	}
	defer data.Close(ctx)

	for data.Next(ctx) {
		var rowCat *model.Tracker
		err = data.Decode(&rowCat)
		if err != nil {
			return
		}
		res = append(res, rowCat)
	}
	return
}

func (r *MongoClient) FindById(ctx context.Context, id string) (res *model.Tracker, err error) {
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

func (r *MongoClient) FindByVehicleId(ctx context.Context, vehicleId string) (res *model.Tracker, err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"vehicle_id": vehicleId}
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

func (r *MongoClient) Create(ctx context.Context, data *model.Tracker) (err error) {
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

func (r *MongoClient) Update(ctx context.Context, id string, data *model.Tracker) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"vehicle_id": id}
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

func (r *MongoClient) UpdateLocation(ctx context.Context, vehicleId string, data *model.Tracker) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"vehicle_id": vehicleId}
	update := bson.M{"$push": bson.M{"location": bson.M{"latitude": data.Location[0].Latitude, "longitude": data.Location[0].Longitude}}}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return
	}
	return nil
}

func (r *MongoClient) Delete(ctx context.Context, id string) (err error) {
	coll := r.conn.Database(r.db).Collection(collection)
	filter := bson.M{"vehicle_id": id}
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
