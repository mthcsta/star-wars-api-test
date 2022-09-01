package service

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	filmService FilmService = FilmService{}
)

func TestGetFilmsObjectIDById(t *testing.T) {
	ListObjectID := filmService.GetFilmsObjectIDById([]int{1, 2, 3})
	if reflect.TypeOf(ListObjectID[0]) != reflect.TypeOf(primitive.NewObjectID()) {
		t.Errorf("returned type %s", reflect.TypeOf(ListObjectID[0]))
	}
}
func TestCreateFilmAndGetObjectIDByIdRef(t *testing.T) {
	objectId := filmService.CreateFilmAndGetObjectIDByIdRef(1)
	if reflect.TypeOf(objectId) != reflect.TypeOf(primitive.NewObjectID()) {
		t.Errorf("returned type %s", reflect.TypeOf(objectId))
	}
}
