package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddPlanet struct {
	Name    string               `json:"name" bson:"name" example:"Matheus"`
	Climate string               `json:"climate" bson:"climate" example:"arid"`
	Terrain string               `json:"terrain" bson:"terrain" example:"desert"`
	Films   []primitive.ObjectID `json:"-" bson:"films"`
}

type Planet struct {
	Id      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Climate string             `json:"climate" bson:"climate"`
	Terrain string             `json:"terrain" bson:"terrain"`
	Films   []Film             `json:"films" bson:"films"`
}

type FilterPlanet struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty" form:"id"`
	Name    string             `json:"name" bson:"name,omitempty" form:"name"`
	Climate string             `json:"climate" bson:"climate,omitempty" form:"climate"`
	Terrain string             `json:"terrain" bson:"terrain,omitempty" form:"terrain"`
}
