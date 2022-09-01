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

type FilmDAO struct{}

var filmCollection *mongo.Collection = database.GetCollection(database.DB, "films")

func (f *FilmDAO) GetByIdRef(ctx context.Context, id int) (model.Film, error) {
	var film model.Film
	err := filmCollection.FindOne(ctx, bson.M{"id_ref": id}).Decode(&film)
	return film, err
}

func (f *FilmDAO) GetById(ctx context.Context, id primitive.ObjectID) (model.Film, error) {
	var film model.Film
	err := filmCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&film)
	return film, err
}

func (p *FilmDAO) GetAll(ctx context.Context, title string) *[]model.Film {
	var films []model.Film
	cursor, err := filmCollection.Find(ctx, bson.M{"title": bson.M{"$regex": title}})
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
		var film model.Film
		bm, _ := bson.Marshal(result)
		bson.Unmarshal(bm, &film)
		films = append(films, film)
	}
	return &films
}

func (p *FilmDAO) InsertOne(ctx context.Context, film model.AddFilm) *mongo.InsertOneResult {
	result, err := filmCollection.InsertOne(ctx, &film)
	if err != nil {
		log.Panic(err)
	}
	return result
}

// func (p *FilmDAO) Remove(ctx context.Context, id primitive.ObjectID) *mongo.DeleteResult {
// 	result, err := planetCollection.DeleteOne(ctx, bson.M{"_id": id})
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return result
// }
