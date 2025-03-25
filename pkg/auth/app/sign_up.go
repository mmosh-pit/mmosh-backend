package auth

import (
	"github.com/gagliardetto/solana-go"
	"github.com/google/uuid"
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	utils "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	commonUtils "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
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

	wallet := solana.NewWallet()

	password, err := utils.EncryptPassword(params.Password)

	if err != nil {
		return nil, err
	}

	privateKey := commonUtils.EncryptPrivateKey(wallet.PrivateKey.String())

	uuid, _ := uuid.NewRandom()

	id := primitive.NewObjectID()

	user := &authDomain.User{
		ID:         &id,
		Name:       params.Name,
		Email:      params.Email,
		Bsky:       authDomain.BlueskyData{},
		Address:    wallet.PublicKey().String(),
		Password:   password,
		Sessions:   []string{*token},
		ReferredBy: "",
		PrivateKey: privateKey,
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
