package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddFilm struct {
	IdReference int       `json:"id_reference" bson:"id_ref"`
	Title       string    `json:"title" bson:"title"`
	EpisodeId   int       `json:"episode_id" bson:"episodeId"`
	Director    string    `json:"director" bson:"director"`
	ReleaseDate time.Time `json:"release_date" bson:"releaseDate"`
}

type Film struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	IdReference int                `json:"id_reference" bson:"id_ref"`
	Title       string             `json:"title" bson:"title"`
	EpisodeId   int                `json:"episode_id" bson:"episodeId"`
	Director    string             `json:"director" bson:"director"`
	ReleaseDate time.Time          `json:"release_date" bson:"releaseDate"`
}

type FilterFilm struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	EpisodeId   int                `json:"episode_id" bson:"episodeId,omitempty"`
	Director    string             `json:"director" bson:"director,omitempty"`
	ReleaseDate time.Time          `json:"release_date" bson:"releaseDate,omitempty"`
}

// type FilterFilm struct {
// 	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty" form:"id"`
// 	Name    string             `json:"name" bson:"name,omitempty" form:"name"`
// 	Climate string             `json:"climate" bson:"climate,omitempty" form:"climate"`
// 	Terrain string             `json:"terrain" bson:"terrain,omitempty" form:"terrain"`
// }
