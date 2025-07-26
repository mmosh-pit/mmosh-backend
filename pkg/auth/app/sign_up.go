package auth

import (
	"github.com/google/uuid"
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	utils "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	commonApp "github.com/mmosh-pit/mmosh_backend/pkg/common/app"
	commonDomain "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpResponse struct {
	Token *string          `json:"token"`
	User  *authDomain.User `json:"user"`
}

func SignUp(params *authDomain.SignUpParams) (*SignUpResponse, error) {
	existingUser, err := authDb.GetUserByEmail(params.Email)

	if err == nil && existingUser.ID != nil {
		return nil, authDomain.ErrUserAlreadyExists
	}

	existingCode := authDb.GetTemporalCode(params.Code)

	if existingCode == nil {
		return nil, commonDomain.InvalidCodeErr
	}

	if existingCode.Email != params.Email {
		return nil, commonDomain.InvalidCodeErr
	}

	authDb.DeleteTemporalCode(existingCode.Code)

	token, err := utils.GenerateSessionToken([]string{})

	if err != nil {
		return nil, err
	}

	address, err := CreateWallet(params.Email)

	if err != nil {
		return nil, err
	}

	password, err := utils.EncryptPassword(params.Password)

	if err != nil {
		return nil, err
	}

	uuid, _ := uuid.NewRandom()

	id := primitive.NewObjectID()

	bot := params.FromBot

	if bot == "" {
		bot = "KIN"
	}

	user := &authDomain.User{
		ID:         &id,
		Name:       params.Name,
		Email:      params.Email,
		Password:   password,
		Sessions:   []string{*token},
		ReferredBy: "",
		UUID:       uuid.String(),
		Wallet:     address,
		GuestData: authDomain.GuestUserData{
			Picture: "https://storage.googleapis.com/mmosh-assets/default.png",
			Name:    params.Name,
		},
		FromBot: bot,
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

	chatDb.SetDefaultChat(user, bot)

	if bot == "FDN" {
		go commonApp.SendKartraNotification("full_disclosure_bot", params.Name, "", params.Email)
	} else {
		go commonApp.SendKartraNotification("Kinship_Bots_Sign_Up", params.Name, "", params.Email)
	}

	return response, nil
}
