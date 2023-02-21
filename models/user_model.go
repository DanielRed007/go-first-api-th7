package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auto struct {
    Id primitive.ObjectID `json:"id,omitempty" bson:"_id"`
    Marca string             `json:"marca,omitempty"`
    Modelo string             `json:"modelo,omitempty"`
    Año int            `json:"año,omitempty"`
	Color string             `json:"color,omitempty"`
}