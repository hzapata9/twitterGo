package bd

import (
	"context"
	"twitterGo/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(usersCollection models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("Usuarios")

	usersCollection.Password, _ = EncriptarPassword(usersCollection.Password)

	result, err := col.InsertOne(ctx, usersCollection)
	if err != nil {
		return "", false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil
}
