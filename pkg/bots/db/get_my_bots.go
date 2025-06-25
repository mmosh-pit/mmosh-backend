package bots

import (
	"log"

	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMyBots(profilenft string) (*[]bots.Bot, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	db := client.Database(databaseName) // Replace with your database name
	collection := db.Collection("mmosh-app-project")

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "status", Value: "active"}, {Key: "creator", Value: profilenft}}}}
	limitStage := bson.D{{Key: "$limit", Value: 100}}

	pipeline := mongo.Pipeline{
		matchStage, // Your dynamic match stage
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "mmosh-app-project-coins"},
			{Key: "localField", Value: "key"},
			{Key: "foreignField", Value: "projectkey"},
			{Key: "as", Value: "coins"},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "mmosh-app-project-community"},
			{Key: "localField", Value: "key"},
			{Key: "foreignField", Value: "projectkey"},
			{Key: "as", Value: "community"},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "mmosh-app-project-profiles"},
			{Key: "localField", Value: "key"},
			{Key: "foreignField", Value: "projectkey"},
			{Key: "as", Value: "profiles"},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "mmosh-app-project-tokenomics"},
			{Key: "localField", Value: "key"},
			{Key: "foreignField", Value: "projectkey"},
			{Key: "as", Value: "tokenomics"},
		}}},
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "mmosh-app-project-pass"},
			{Key: "localField", Value: "key"},
			{Key: "foreignField", Value: "projectkey"},
			{Key: "as", Value: "pass"},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "name", Value: 1},
			{Key: "symbol", Value: 1},
			{Key: "desc", Value: 1},
			{Key: "key", Value: 1},
			{Key: "image", Value: 1},
			{Key: "price", Value: 1},
			{Key: "coins", Value: "$coins"},
			{Key: "community", Value: "$community"},
			{Key: "profiles", Value: "$profiles"},
			{Key: "tokenomics", Value: "$tokenomics"},
			{Key: "pass", Value: "$pass"},
		}}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "created_date", Value: -1}}}},
		limitStage,
	}

	cursor, err := collection.Aggregate(*ctx, pipeline)
	if err != nil {
		log.Printf("Error in aggregation: %v\n", err)
		return nil, err
	}
	defer cursor.Close(*ctx)

	var results []bots.Bot

	if err = cursor.All(*ctx, &results); err != nil {
		log.Printf("Error decoding results: %v\n", err)
		return nil, err
	}

	return &results, nil
}
