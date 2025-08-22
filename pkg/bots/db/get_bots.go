package bots

import (
	"log"

	botsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBots(search, username string, page int64, isWizard bool) []botsDomain.Bot {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-project")

	filter := bson.D{{}}

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

	if !isWizard {
		filter = bson.D{{
			Key: "$or",
			Value: []any{
				bson.D{{
					Key: "privacy",
					Value: bson.D{{
						Key:   "$exists",
						Value: false,
					}},
				}},

				bson.D{{
					Key:   "privacy",
					Value: "public",
				}},

				bson.D{{
					Key: "$and",
					Value: []any{
						bson.D{{
							Key:   "creatorUsername",
							Value: username,
						}},

						bson.D{{
							Key: "privacy",
							Value: bson.D{{
								Key:   "$exists",
								Value: false,
							}},
						}},
					},
				}},
			},
		}}
	}

	if search != "" {
		filter = bson.D{{
			Key: "$and",
			Value: []any{
				bson.D{{
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

						bson.D{{
							Key: "$and",
							Value: []any{
								bson.D{{
									Key:   "code",
									Value: search,
								}},

								bson.D{{
									Key:   "privacy",
									Value: "secret",
								}},
							},
						}},
					},
				}},
				// bson.D{{
				// 	Key: "$and",
				// 	Value: []any{
				// 		bson.D{{
				// 			Key: "creatorUsername",
				// 			Value: bson.D{{
				// 				Key:   "$ne",
				// 				Value: username,
				// 			}},
				// 		}},
				//
				// 		bson.D{{
				// 			Key:   "privacy",
				// 			Value: "hidden",
				// 		}},
				// 	},
				// }},

				bson.D{{
					Key: "$and",
					Value: []any{
						bson.D{{
							Key: "creatorUsername",
							Value: bson.D{{
								Key:   "$ne",
								Value: username,
							}},
						}},

						bson.D{{
							Key:   "privacy",
							Value: "private",
						}},

						bson.D{{
							Key: "privacy",
							Value: bson.D{{
								Key:   "$exists",
								Value: true,
							}},
						}},
					},
				}},
			},
		}}
	}

	var result = []botsDomain.Bot{}

	res, err := collection.Find(*ctx, filter, opts)

	if err != nil {
		log.Printf("[GET BOTS] Got error here: %v\n", err)
		return result
	}

	for res.Next(*ctx) {
		var bot botsDomain.Bot

		if err := res.Decode(&bot); err != nil {
			log.Printf("[GET BOTS] Error decoding bot: %v\n", err)
			continue
		}

		result = append(result, bot)
	}

	return result
}
