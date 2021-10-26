package bd

import (
	"context"
	"time"

	"github.com/kevin1892/twetter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twetter")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPAssword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID(primitive.ObjectID)

	return ObjID.String(), true, nil
}
