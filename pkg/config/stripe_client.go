package config

import (
	"github.com/stripe/stripe-go/v82"
)

var stripeClient *stripe.Client

func InitializeStripe() {
	stripeApiKey := GetStripeVariable()
	stripeClient = stripe.NewClient(stripeApiKey, stripe.WithBackends(
		stripe.NewBackendsWithConfig(&stripe.BackendConfig{
			LeveledLogger: &stripe.LeveledLogger{Level: stripe.LevelNull},
		}),
	))
}

func GetStripeClient() *stripe.Client {
	return stripeClient
}
