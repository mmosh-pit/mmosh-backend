package auth

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrSomethingWentWrong = errors.New("Something Went Wrong")
var ErrWalletAlreadyExists = errors.New("wallet-exists")
var ErrUserAlreadyExists = errors.New("user-exists")
var ErrDataAlreadyExists = errors.New("data-already-exists")

type OnboardingStepParams struct {
	Step int `json:"step"`
}

type LoginParams struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type SignUpParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
}

type GuestUserData struct {
	Picture  string `json:"picture"`
	Banner   string `jsos:"banner"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Website  string `json:"website"`
	Pronouns string `json:"pronouns"`
	Bio      string `json:"bio"`
}

type AddReferrerParams struct {
	User string `json:"user"`
}

type User struct {
	ID             *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	UUID           string              `json:"uuid" bson:"uuid"`
	Name           string              `bson:"name" json:"name"`
	Email          string              `bson:"email" json:"email"`
	Password       string              `bson:"password" json:"password"`
	TelegramId     int                 `bson:"telegramId" json:"telegramId"`
	GuestData      GuestUserData       `bson:"guest_data" json:"guest_data"`
	Sessions       []string            `bson:"sessions" json:"sessions"`
	Bsky           BlueskyData         `bson:"bsky" json:"bsky"`
	Subscription   UserSubscription    `bson:"subscription" json:"subscription"`
	Wallet         string              `json:"wallet" bson:"wallet"`
	ReferredBy     string              `json:"referred_by" bson:"referred_by"`
	OnboardingStep int                 `json:"onboarding_step" bson:"onboarding_step"`
	CreatedAt      time.Time           `bson:"created_at"`
}

type UserSubscription struct {
	ProductId        string `bson:"product_id" json:"product_id"`
	PurchaseToken    string `bson:"purchase_token" json:"purchase_token"`
	SubscriptionId   string `bson:"subscription_id" json:"subscription_id"`
	SubscriptionTier int    `bson:"subscription_tier" json:"subscription_tier"`
	ExpiresAt        int64  `bson:"expires_at" json:"expires_at"`
	Platform         string `bson:"platform" json:"platform"`
	ChangedPlan      bool   `bson:"changed_plan" json:"changed_plan"`
}

type BlueskyData struct {
	Id           string `bson:"id" json:"id"`
	Handle       string `bson:"handle" json:"handle"`
	Token        string `bson:"token" json:"token"`
	RefreshToken string `bson:"refreshToken" json:"refreshToken"`
}

type VerificationData struct {
	Email string `bson:"email" json:"email"`
	Code  int    `bson:"code" json:"code"`
}

type KeyPair struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type WalletResponse struct {
	Address    string   `json:"address"`
	KeyPackage []string `json:"key_package"`
}

type Wallet struct {
	Address   string    `json:"address" bson:"address"`
	Private   string    `json:"private" bson:"private"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
