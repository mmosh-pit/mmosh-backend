package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AppTheme struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	CodeName        string             `json:"code_name" bson:"code_name"`
	BackgroundColor string             `json:"background_color" bson:"background_color"`
	PrimaryColor    string             `json:"primary_color" bson:"primary_color"`
	SecondaryColor  string             `json:"secondary_color" bson:"secondary_color"`
	Logo            string             `json:"logo" bson:"logo"`
}
