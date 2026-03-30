package auth

import (
	"errors"
	"time"
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
	Name   string `json:"name"`
	Email  string `json:"email"`
	Reason string `json:"reason"`
}

type AddReferrerParams struct {
	User string `json:"user"`
}

type WebsiteLink struct {
	Link  string `json:"link"`
	Order int    `json:"order"`
}

type User struct {
	ID          string `json:"id"`
	UUID        string `json:"uuid"`
	Picture     string `json:"picture"`
	Banner      string `json:"banner"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	Websites    []WebsiteLink `json:"websites"`
	Bio         string `json:"bio"`
	Challenges  string `json:"challenges"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	Telegram    TelegramUserData `json:"telegram"`
	Sessions    []string `json:"sessions,omitempty"`
	Bluesky     BlueskyUserData `json:"bluesky"`
	Subscription UserSubscription `json:"subscription"`
	Wallet      string `json:"wallet,omitempty"`
	ReferredBy  string `json:"referred_by,omitempty"`
	OnboardingStep int `json:"onboarding_step,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	LastLogin   time.Time `json:"lastLogin"`
	ProfileNFT  string `json:"profilenft,omitempty"`
	Role        string `json:"role,omitempty"`
	FromBot     string `json:"from_bot,omitempty"`
	Deactivated bool   `json:"deactivated"`

	Seniority       int    `json:"seniority"`
	Symbol          string `json:"symbol"`
	Link            string `json:"link"`
	Following       int    `json:"following"`
	Follower        int    `json:"follower"`
	ConnectionNFT   string `json:"connectionnft"`
	ConnectionBadge string `json:"connectionbadge"`
	Connection      int    `json:"connection"`
	IsPrivate       bool   `json:"isprivate"`
	Request         bool   `json:"request"`
}

type UserSubscription struct {
	ProductId        string `json:"product_id"`
	SubProductId     string `json:"sub_product_id"`
	PurchaseToken    string `json:"purchase_token"`
	SubscriptionId   string `json:"subscription_id"`
	SubscriptionTier int    `json:"subscription_tier"`
	ExpiresAt        int64  `json:"expires_at"`
	Platform         string `json:"platform"`
	ChangedPlan      bool   `json:"changed_plan"`
}

type BlueskyUserData struct {
	Handle   string `json:"handle"`
	Password string `json:"password"`
}

type TelegramUserData struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	Username  string `json:"username"`
}

type BlueskyData struct {
	Id           string `json:"id"`
	Handle       string `json:"handle"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type VerificationData struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
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
	Address   string    `json:"address"`
	Private   string    `json:"private"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddEarlyAccessParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
