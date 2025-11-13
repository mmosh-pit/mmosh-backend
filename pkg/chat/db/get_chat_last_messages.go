package chat

import (
	"log"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetChatLastMessages(chatId *primitive.ObjectID) []chat.Message {

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("chats")

	opts := options.Find().SetLimit(20).SetSort(bson.D{{
		Key:   "created_at",
		Value: -1,
	}})

	var result = []chat.Message{}

	res, err := collection.Find(*ctx, bson.D{{
		Key:   "_id",
		Value: chatId,
	}}, opts)

	if err != nil {
		log.Printf("[GET CHAT LAST MESSAGES] Got error here: %v\n", err)
		return result
	}

	for res.Next(*ctx) {
		var bot chat.Message

		if err := res.Decode(&bot); err != nil {
			log.Printf("[GET CHAT LAST MESSAGES] Error decoding message: %v\n", err)
			continue
		}

		result = append(result, bot)
	}

	return result

}
