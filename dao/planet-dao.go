package dao

import (
	"context"
	"log"

	"github.com/mthcsta/star-wars-api-test/database"
	"github.com/mthcsta/star-wars-api-test/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanetDAO struct{}

var planetCollection *mongo.Collection = database.GetCollection(database.DB, "planets")

func (p *PlanetDAO) GetAll(ctx context.Context, filterPlanet model.FilterPlanet) *[]model.Planet {
	var planets []model.Planet
	var bso *bson.M
	marshal, _ := bson.Marshal(filterPlanet)
	bson.Unmarshal(marshal, &bso)
	cursor, err := planetCollection.Find(ctx, *bso)
	defer cursor.Close(ctx)

	if err != nil {
		log.Panic(err)
	}

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			break
		}
		var planet model.Planet
		bm, _ := bson.Marshal(result)
		bson.Unmarshal(bm, &planet)
		planets = append(planets, planet)
	}
	return &planets
}

func (p *PlanetDAO) InsertOne(ctx context.Context, planet model.AddPlanet) *mongo.InsertOneResult {
	result, err := planetCollection.InsertOne(ctx, &planet)
	if err != nil {
		log.Panic(err)
	}
	return result
}

func (p *PlanetDAO) Remove(ctx context.Context, id primitive.ObjectID) *mongo.DeleteResult {
	result, err := planetCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Panic(err)
	}
	return result
}

func (p *PlanetDAO) Exists(ctx context.Context, name string) bool {
	var planet model.Planet
	err := planetCollection.FindOne(ctx, bson.M{"name": name}).Decode(&planet)
	return err != mongo.ErrNoDocuments
}
