package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateProfileDataParams struct {
	Symbol      string `json:"symbol"`
	Bio         string `json:"bio"`
	DisplayName string `json:"displayName"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	Banner      string `json:"banner"`
}

func UpdateProfileData(params UpdateProfileDataParams, userId string) error {

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	profileData := authDomain.Profile{
		Name:        params.Name,
		LastName:    params.LastName,
		DisplayName: params.DisplayName,
		Username:    params.Username,
		Bio:         params.Bio,
		Image:       params.Image,
		Symbol:      params.Symbol,
		IsPrivate:   false,
	}

	err = authDb.UpdateProfileData(profileData, userIdBson)

	authDb.UpdateUserOnboardingStatus(&userIdBson, 4)

	return err
}
