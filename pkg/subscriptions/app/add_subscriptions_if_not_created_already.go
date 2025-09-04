package subscriptions

import (
	"log"
	"strings"

	subscriptionsDb "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionProductIds = []string{"guest", "enjoyer", "creator"}

var benefits = [][]string{
	{
		"Can only interact with Public Bots",
		"No access to Bot Studio, Offers or Bot Subscriptions",
	},

	{
		"Revenue Distribution",
		"Up to 3 Personal Bots",
		"Up to 3 Community Bots",
	},

	{
		"Revenue Distribution",
		"Up to 3 Personal Bots",
	},
}

func AddSubscriptionsIfNotCreatedAlready() {
	for i, item := range subscriptionProductIds {
		_, err := subscriptionsDb.GetSubscriptionByProductId(item)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				err = subscriptionsDb.AddSubscription(&subscriptionsDomain.Subscription{
					ID:        primitive.NewObjectID(),
					Tier:      i + 1,
					ProductId: item,
					Name:      strings.Title(item),
					Benefits:  benefits[i],
				})

				if err != nil {
					log.Printf("Could not create document with product id: %v, %v\n", item, err)
				}
			}
		}
	}
}
