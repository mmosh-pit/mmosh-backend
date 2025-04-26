package auth

import (
	"github.com/google/uuid"
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	utils "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpResponse struct {
	Token *string          `json:"token"`
	User  *authDomain.User `json:"user"`
}

func SignUp(params *authDomain.SignUpParams) (*SignUpResponse, error) {
	existingCode := authDb.GetTemporalCode(params.Code)

	if existingCode == nil {
		return nil, common.InvalidCodeErr
	}

	if existingCode.Email != params.Email {
		return nil, common.InvalidCodeErr
	}

	authDb.DeleteTemporalCode(existingCode.Code)

	token, err := utils.GenerateSessionToken([]string{})

	if err != nil {
		return nil, err
	}

	err = CreateWallet(params.Email)

	if err != nil {
		return nil, err
	}

	password, err := utils.EncryptPassword(params.Password)

	if err != nil {
		return nil, err
	}

	uuid, _ := uuid.NewRandom()

	id := primitive.NewObjectID()

	user := &authDomain.User{
		ID:         &id,
		Name:       params.Name,
		Email:      params.Email,
		Bsky:       authDomain.BlueskyData{},
		Password:   password,
		Sessions:   []string{*token},
		ReferredBy: "",
		UUID:       uuid.String(),
	}

	err = authDb.CreateUser(user)

	if err != nil {
		return nil, err
	}

	user.Sessions = []string{}
	user.Password = ""

	response := &SignUpResponse{
		Token: token,
		User:  user,
	}

	chatDb.CreateChat(user)

	return response, nil
}
