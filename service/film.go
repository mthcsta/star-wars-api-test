package service

import (
	"context"
	"sync"
	"time"

	"github.com/mthcsta/star-wars-api-test/apis"
	"github.com/mthcsta/star-wars-api-test/dao"
	"github.com/mthcsta/star-wars-api-test/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilmService struct {
	filmDAO dao.FilmDAO
}

func (f *FilmService) GetFilmsObjectIDById(filmsId []int) []primitive.ObjectID {
	var filmsObjectID []primitive.ObjectID
	var wg sync.WaitGroup
	for _, filmId := range filmsId {
		wg.Add(1)
		filmId := filmId
		go func() {
			defer wg.Done()
			objectID := f.CreateFilmAndGetObjectIDByIdRef(filmId)
			filmsObjectID = append(filmsObjectID, objectID)
		}()
	}
	wg.Wait()

	return filmsObjectID
}

func (f *FilmService) CreateFilmAndGetObjectIDByIdRef(id int) primitive.ObjectID {
	film, err := f.filmDAO.GetByIdRef(context.Background(), id)
	if err == mongo.ErrNoDocuments {
		json := apis.GetFilmById(id)
		releaseDate, _ := time.Parse("2006-01-02", json.ReleaseDate)

		newFilm := model.AddFilm{
			IdReference: id,
			Title:       json.Title,
			EpisodeId:   json.EpisodeId,
			Director:    json.Director,
			ReleaseDate: releaseDate,
		}

		result := f.filmDAO.InsertOne(context.Background(), newFilm)
		return result.InsertedID.(primitive.ObjectID)
	}
	return film.Id
}
