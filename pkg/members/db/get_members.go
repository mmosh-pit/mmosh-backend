package members

import (
	"log"
	"regexp"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMembers(page int64, search string, userId primitive.ObjectID) []auth.User {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	var result = []auth.User{}

	opts := options.Find().SetSkip(page * 20).SetLimit(20).SetProjection(
		bson.M{"sessions": 0, "password": 0, "bluesky.password": 0},
	).SetSort(bson.D{{
		Key:   "profile.seniority",
		Value: -1,
	}})

	filter := bson.D{{
		Key: "_id",
		Value: bson.D{{
			Key:   "$ne",
			Value: userId,
		}},
	},
		{
			Key: "profilenft",
			Value: bson.D{{
				Key:   "$exists",
				Value: true,
			}},
		},
	}

	if search != "" {
		re := regexp.MustCompile(`/[/\-\\^$*+?.()|[\]{}]/g`)

		searchText := re.ReplaceAllString(search, "")

		filter = bson.D{{
			Key: "$and",
			Value: []any{
				bson.D{{
					Key: "_id",
					Value: bson.D{{
						Key:   "$ne",
						Value: userId,
					}},
				},
					{
						Key: "profilenft",
						Value: bson.D{{
							Key:   "$exists",
							Value: true,
						}},
					},
				},

				bson.D{{
					Key: "$or",
					Value: []any{
						bson.D{{
							Key:   "profile.username",
							Value: primitive.Regex{Pattern: searchText, Options: "i"},
						}},

						bson.D{{
							Key:   "profile.name",
							Value: primitive.Regex{Pattern: searchText, Options: "i"},
						}},

						bson.D{{
							Key:   "guest_data.name",
							Value: primitive.Regex{Pattern: searchText, Options: "i"},
						}},

						bson.D{{
							Key:   "guest_data.username",
							Value: primitive.Regex{Pattern: searchText, Options: "i"},
						}},
					},
				}},
			},
		},
		}
	}

	res, err := collection.Find(*ctx, filter, opts)

	if err != nil {
		log.Printf("Got error retrieving members: %v\n", err)
		return result
	}

	for res.Next(*ctx) {
		var user auth.User

		if err := res.Decode(&user); err != nil {
			log.Printf("Error decoding user: %v\n", err)
			continue
		}

		result = append(result, user)
	}

	return result
}
