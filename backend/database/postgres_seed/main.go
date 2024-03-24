package main

import (
	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/logger"
)

func main() {
	c, err := client.GetPsqlClient()
	if err != nil {
		panic(err)
	}

	um := models.UserModel{
		UserId:           "auth0|65ffab5c004e8d1620d06a64",
		Name:             "Test User",
		IdentityProvider: "auth0",
	}

	user, err := database.NewUser(c).Create(&um)
	if err != nil {
		logger.Error("Error inserting user", map[string]interface{}{"err": err})
		panic(err)
	}
	logger.Info("Inserted user", map[string]interface{}{"user": user})

	om := models.OrganisationModel{
		Name:  "Test Org",
		Owner: "auth0|65ffab5c004e8d1620d06a64",
	}
	org, err := database.NewOrganization(c).Create(&om)
	if err != nil {
		logger.Error("Error inserting org", map[string]interface{}{"err": err})
		panic(err)
	}
	logger.Info("Created organisation and membership", map[string]interface{}{"org": org})

}
