package db

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(page int64, search string) *[]auth.User {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	var result []auth.User

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{}

	opts := options.Find().SetSkip(page * 20).SetLimit(20).SetProjection(
		bson.M{
			"_id":             1,
			"name":            1,
			"symbol":          1,
			"desc":            1,
			"key":             1,
			"image":           1,
			"creatorUsername": 1,
			"privacy":         1,
			"system_prompt":   1,
			"type":            1,
		},
	).SetSort(bson.D{{
		Key:   "profile.seniority",
		Value: -1,
	}})

	if search != "" {
		filter = bson.D{{
			Key: "$or",
			Value: []any{
				bson.D{{
					Key: "name",
					Value: primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				}},

				bson.D{{
					Key: "symbol",
					Value: primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				}},

				bson.D{{
					Key: "desc",
					Value: primitive.Regex{
						Pattern: search,
						Options: "i",
					},
				}},
			},
		}}
	}

	res, err := collection.Find(*ctx, filter, opts)

	if err != nil {
		log.Printf("[ADMIN/GET ALL USERS] Got error retrieving all users: %v\n", err)
		return &result
	}

	for res.Next(*ctx) {
		var user auth.User

		if err := res.Decode(&user); err != nil {
			log.Printf("[ADMIN/GET ALL USERS] could not decode user: %v\n", err)
			continue
		}

		result = append(result, user)
	}

	return &result
}
