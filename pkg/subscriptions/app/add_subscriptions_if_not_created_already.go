package subscriptions

import (
	"log"
	"strings"

	subscriptionsDb "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
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
		existing, err := subscriptionsDb.GetSubscriptionByProductId(item)

		if err != nil {
			log.Printf("Could not check subscription with product id: %v, %v\n", item, err)
			continue
		}

		if existing == nil {
			err = subscriptionsDb.AddSubscription(&subscriptionsDomain.Subscription{
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
