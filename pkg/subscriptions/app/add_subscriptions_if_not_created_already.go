package subscriptions

import (
	"log"
	"strings"

	subscriptionsDb "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionProductIds = []string{
	"guest",
	// "enjoyer",
	// "creator",
	"member",
}

var benefits = [][]string{
	{
		"Limited AI access",
		"No connections",
		"Basic communities",
		"No signals",
		"No referral rewards",
	},

	// {
	// 	"Revenue Distribution",
	// 	"Up to 3 Personal Bots",
	// 	"Up to 3 Community Bots",
	// },
	//
	// {
	// 	"Revenue Distribution",
	// 	"Up to 3 Personal Bots",
	// },

	{
		"Unlimited AI access",
		"Unlimited connections",
		"Member-only communities",
		"Unlimited Signals (AI-mediated DMs)",
		"Instant cash referral rewards",
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
