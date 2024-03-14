package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDb(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()
}

func Ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

func InsertOne(client *mongo.Client, ctx context.Context, database, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(database).Collection(col)

	result, err := collection.InsertOne(ctx, doc)

	return result, err

}

func InsertMany(client *mongo.Client, ctx context.Context, database, col string, doc []interface{}) (*mongo.InsertManyResult, error) {

	collection := client.Database(database).Collection(col)

	result, err := collection.InsertMany(ctx, doc)

	return result, err
}

func Query(client *mongo.Client, ctx context.Context, database, col string, filter, option interface{}) (*mongo.Cursor, error) {
	collection := client.Database(database).Collection(col)

	result, err := collection.Find(ctx, filter, options.Find().SetProjection(option))

	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, database, col string, filter, updateData interface{}) (*mongo.UpdateResult, error) {

	collection := client.Database(database).Collection(col)

	result, err := collection.UpdateOne(ctx, filter, updateData)

	return result, err
}

func UpdateMany(client *mongo.Client, ctx context.Context, database, col string, filter, updateData []interface{}) (*mongo.UpdateResult, error) {

	collection := client.Database(database).Collection(col)

	result, err := collection.UpdateOne(ctx, filter, updateData)

	return result, err
}

func DeleteOne(client *mongo.Client, ctx context.Context, database, col string, filter interface{}) (*mongo.DeleteResult, error) {

	collection := client.Database(database).Collection(col)

	result, err := collection.DeleteOne(ctx, filter)

	return result, err
}

func DeleteMany(client *mongo.Client, ctx context.Context, database, col string, filter interface{}) (*mongo.DeleteResult, error) {

	collection := client.Database(database).Collection(col)

	result, err := collection.DeleteMany(ctx, filter)

	return result, err
}
