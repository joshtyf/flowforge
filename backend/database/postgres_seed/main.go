package main

import (
	"fmt"
	"os"

	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
)

func main() {
	logger := logger.NewServerLog(os.Stdout)
	c, err := client.GetPsqlClient()
	if err != nil {
		panic(err)
	}

	um1 := models.UserModel{
		UserId:           "auth0|65ffab5c004e8d1620d06a64",
		Name:             "Test User 1",
		IdentityProvider: "auth0",
	}

	um2 := models.UserModel{
		UserId:           "auth0|66010ad5095367b237799680",
		Name:             "Test User 2",
		IdentityProvider: "auth0",
	}

	um3 := models.UserModel{
		UserId:           "auth0|65e9dabff2dab546ed0c231e",
		Name:             "Test User 3",
		IdentityProvider: "auth0",
	}

	users := [...]models.UserModel{um1, um2, um3}
	for i := 0; i < 3; i++ {
		user, err := database.NewUser(c).Create(&users[i])
		if err != nil {
			logger.Error(fmt.Sprintf("unable to insert user: %s", err))
			panic(err)
		}
		logger.Info(fmt.Sprintf("inserted user: %v", user))
	}

	om := models.OrganizationModel{
		Name:  "Test Org",
		Owner: "auth0|65ffab5c004e8d1620d06a64",
	}

	org, err := database.NewOrganization(c).Create(&om)
	if err != nil {
		logger.Error(fmt.Sprintf("unable to insert org: %s", err))
		panic(err)
	}
	logger.Info(fmt.Sprintf("inserted org: %v", org))

	_, err = database.NewMembership(c).Create(&models.MembershipModel{UserId: um2.UserId, OrgId: org.OrgId, Role: models.Admin})
	if err != nil {
		logger.Error(fmt.Sprintf("unable to create membership: %s", err))
		panic(err)
	}

	_, err = database.NewMembership(c).Create(&models.MembershipModel{UserId: um3.UserId, OrgId: org.OrgId, Role: models.Member})
	if err != nil {
		logger.Error(fmt.Sprintf("unable to create membership: %s", err))
		panic(err)
	}
}
