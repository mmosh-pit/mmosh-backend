package app

// func GetOnboardURL(userId string) (*string, error) {
// 	ctx, db := config.GetPostgreSqlWriteConnection()
//
// 	var foundUser struct {
// 		PendingStripeAccountId *string `ksql:"pending_stripe_account_id"`
// 	}
// 	if err := db.QueryOne(
// 		ctx,
// 		&foundUser,
// 		fmt.Sprintf(
// 			"SELECT pending_stripe_account_id FROM %s WHERE id = $1;",
// 			commonDomain.INDEX_USER,
// 		),
// 		userId,
// 	); err != nil {
// 		return nil, err
// 	}
//
// 	stripeClient := config.GetStripeClient()
// 	var stripeAccountId *string
//
// 	if foundUser.PendingStripeAccountId == nil {
// 		account, err := stripeClient.V1Accounts.Create(
// 			context.TODO(),
// 			&stripe.AccountCreateParams{
// 				Country: stripe.String("US"),
// 				Type:    stripe.String("express"),
// 			},
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		result, err := db.Exec(
// 			ctx,
// 			fmt.Sprintf(
// 				"UPDATE %s SET updated_at = now(), pending_stripe_account_id = $2 WHERE id = $1;",
// 				commonDomain.INDEX_USER,
// 			),
// 			userId,
// 			account.ID,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		if affectedRows, _ := result.RowsAffected(); affectedRows == 0 {
// 			if _, err := stripeClient.V1Accounts.Delete(
// 				context.TODO(),
// 				account.ID,
// 				nil,
// 			); err != nil {
// 				log.Println("Error in stripeClient.V1Accounts.Delete: ", err)
// 			}
//
// 			return nil, errors.New("user-not-found")
// 		}
//
// 		stripeAccountId = &account.ID
// 	} else {
// 		stripeAccountId = foundUser.PendingStripeAccountId
// 	}
//
// 	log.Println(config.StripeAccountOnboardingRefreshURL, config.StripeAccountOnboardingReturnURL)
// 	response, err := stripeClient.V1AccountLinks.Create(
// 		context.TODO(),
// 		&stripe.AccountLinkCreateParams{
// 			Account:    stripeAccountId,
// 			RefreshURL: &config.StripeAccountOnboardingRefreshURL,
// 			ReturnURL:  &config.StripeAccountOnboardingReturnURL,
// 			Type:       stripe.String("account_onboarding"),
// 		},
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &response.URL, nil
// }
