package bike

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoBike struct {
	ID    primitive.ObjectID `bson:"_id"`
	Model string             `bson:"model"`
}

type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoRepository(client *mongo.Client) (mongoRepository MongoRepository, err error) {
	collection := client.Database("bikes").Collection("bikes")
	ctx := context.Background()

	mongoRepository = MongoRepository{
		client:     client,
		collection: collection,
		ctx:        ctx,
	}

	return
}

func (repo MongoRepository) GetByFilter(filter Filter) (bikes []Bike, err error) {
	bsonFilter := bson.M{}

	if filter.ID != "" {
		bsonFilter["_id"], err = primitive.ObjectIDFromHex(filter.ID)
		if err != nil {
			return
		}
	}

	if filter.Model != "" {
		bsonFilter["model"] = filter.Model
	}

	cursor, err := repo.collection.Find(repo.ctx, bsonFilter)
	if err != nil {
		return
	}

	defer cursor.Close(repo.ctx)

	bikes = []Bike{}

	for cursor.Next(context.TODO()) {
		var bike Bike

		err = cursor.Decode(&bike)
		if err != nil {
			return
		}

		bikes = append(bikes, bike)
	}

	return
}

func (repo MongoRepository) GetByID(bikeID string) (bike Bike, err error) {
	bikes, err := repo.GetByFilter(Filter{ID: bikeID})
	if err != nil {
		return
	}

	if len(bikes) > 1 {
		err = errors.New("found more that one bike")
		return
	}

	if len(bikes) < 1 {
		err = ErrBikeNotFound
		return
	}

	bike = bikes[0]
	return
}

func (repo MongoRepository) Register(bike Bike) (bikeID string, err error) {
	ctx := context.Background()

	document := bikeToMongoBike(bike)

	result, err := repo.collection.InsertOne(ctx, document)
	if err != nil {
		return
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err = errors.New("invalid object id")
	}

	bikeID = objectID.Hex()
	return
}

func (repo MongoRepository) Update(bike Bike) (err error) {
	objectID, err := primitive.ObjectIDFromHex(bike.ID)
	if err != nil {
		return
	}

	update := bson.M{"$set": bson.M{
		"model": bike.Model,
	}}

	_, err = repo.collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	return
}

func (repo MongoRepository) Delete(bikeID string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(bikeID)
	if err != nil {
		return
	}

	_, err = repo.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return
}

func (repo MongoRepository) Start() (err error) {
	err = repo.client.Database("bikes").CreateCollection(context.Background(), "bikes")
	return
}

func bikeToMongoBike(bike Bike) (mbike mongoBike) {
	mbike = mongoBike{
		ID:    primitive.NewObjectID(),
		Model: bike.Model,
	}

	return
}
