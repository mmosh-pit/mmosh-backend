package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type User struct {
	ID         *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Name       string              `bson:"name" json:"name"`
	Email      string              `bson:"email" json:"email"`
	Password   string              `bson:"password" json:"password"`
	Address    string              `bson:"address" json:"address"`
	ReferredBy string              `bson:"referredBy" json:"referredBy"`
	TelegramId int                 `bson:"telegramId" json:"telegramId"`
	PrivateKey string              `bson:"privateKey" json:"privateKey"`
	Sessions   []string            `bson:"sessions" json:"sessions"`
	Bsky       BlueskyData         `bson:"bsky" json:"bsky"`
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
