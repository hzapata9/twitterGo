package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email,omitempty"`
	Password        string             `bson:"password,omitempty"`
	Avatar          string             `bson:"avatar,omitempty"`
	Banner          string             `bson:"banner,omitempty"`
	Biografia       string             `bson:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioweb,omitempty"`
}
