package subscriptions

import (
	"log"
	"strings"

	subscriptionsDb "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionProductIds = []string{"intensive", "immersive", "expansive"}

var agents = []string{"5", "25", "0"}

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
					Agents:    agents[i],
				})

				if err != nil {
					log.Printf("Could not create document with product id: %v, %v\n", item, err)
				}
			}
		}
	}
}
