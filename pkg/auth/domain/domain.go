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
var ErrInvalidBluesky = errors.New("invalid-bluesky")
var ErrEarlyAlreadyRegistered = errors.New("already-registered")
var ErrUserNotExists = errors.New("user-not-exists")

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
	FromBot  string `json:"from_bot"`
}

type AccountDeletionRequest struct {
	Name   string `json:"name" bson:"name"`
	Email  string `json:"email" bson:"email"`
	Reason string `json:"reason" bson:"reason"`
}

type GuestUserData struct {
	Picture     string `json:"picture"`
	Banner      string `jsos:"banner"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	Website     string `json:"website"`
	Bio         string `json:"bio"`
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
	Telegram       TelegramUserData    `bson:"telegram" json:"telegram"`
	GuestData      GuestUserData       `bson:"guest_data" json:"guest_data"`
	Sessions       []string            `bson:"sessions" json:"sessions"`
	Bluesky        BlueskyUserData     `bson:"bluesky" json:"bluesky"`
	Subscription   UserSubscription    `bson:"subscription" json:"subscription"`
	Wallet         string              `json:"wallet" bson:"wallet"`
	ReferredBy     string              `json:"referred_by" bson:"referred_by"`
	OnboardingStep int                 `json:"onboarding_step" bson:"onboarding_step"`
	CreatedAt      time.Time           `bson:"created_at"`
	Profile        Profile             `json:"profile" bson:"profile"`
	ProfileNFT     string              `json:"profilenft" bson:"profilenft"`
	Role           string              `json:"role" bson:"role"`
	FromBot        string              `bson:"from_bot"`
}

type Profile struct {
	Name            string `json:"name" bson:"name"`
	LastName        string `json:"lastName" bson:"lastName"`
	DisplayName     string `json:"displayName" bson:"displayName"`
	Username        string `json:"username" bson:"username"`
	Bio             string `json:"bio" bson:"bio"`
	Image           string `json:"image" bson:"image"`
	Seniority       int    `json:"seniority" bson:"seniority"`
	Symbol          string `json:"symbol" bson:"symbol"`
	Link            string `json:"link" bson:"link"`
	Following       int    `json:"following" bson:"following"`
	Follower        int    `json:"follower" bson:"follower"`
	ConnectionNFT   string `json:"connectionnft" bson:"connectionnft"`
	ConnectionBadge string `json:"connectionbadge" bson:"connectionbadge"`
	Connection      int    `json:"connection" bson:"connection"`
	IsPrivate       bool   `json:"isprivate" bson:"isprivate"`
	Request         bool   `json:"request" bson:"request"`
}

type UserSubscription struct {
	ProductId        string `bson:"product_id" json:"product_id"`
	SubProductId     string `bson:"sub_product_id" json:"sub_product_id"`
	PurchaseToken    string `bson:"purchase_token" json:"purchase_token"`
	SubscriptionId   string `bson:"subscription_id" json:"subscription_id"`
	SubscriptionTier int    `bson:"subscription_tier" json:"subscription_tier"`
	ExpiresAt        int64  `bson:"expires_at" json:"expires_at"`
	Platform         string `bson:"platform" json:"platform"`
	ChangedPlan      bool   `bson:"changed_plan" json:"changed_plan"`
}

type BlueskyUserData struct {
	Handle   string `bson:"handle" json:"handle"`
	Password string `bson:"password" json:"password"`
}

type TelegramUserData struct {
	Id        int    `bson:"id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	Username  string `bson:"username" json:"username"`
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

type AddEarlyAccessParams struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
