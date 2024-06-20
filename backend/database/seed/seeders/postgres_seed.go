package seeders

import (
	"fmt"
	"os"

	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/helper"
	"github.com/joshtyf/flowforge/src/logger"
)

func SeedPostgres() {
	logger := logger.NewServerLog(os.Stdout)
	c, err := client.GetPsqlClient()
	if err != nil {
		panic(err)
	}

	users, passwords, err := getUsersFromCsv()
	if err != nil {
		panic(err)
	}

	createUsers := os.Getenv("CREATE_USERS")
	if createUsers == "true" {
		logger.Info("Creating users in Auth0")
		helper.CreateUsersInAuth0(users, passwords)
	}
	// Auth0 automatically concatenates the identity provider and user ID with a pipe separator
	for _, user := range users {
		user.UserId = user.IdentityProvider + "|" + user.UserId
	}

	for i := 0; i < len(users); i++ {
		user, err := database.NewUser(c).Create(users[i])
		if err != nil {
			logger.Error(fmt.Sprintf("unable to insert user: %s", err))
			panic(err)
		}
		logger.Info(fmt.Sprintf("inserted user: %v", user))
	}

	orgs, err := getOrgsFromCsv(users)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(orgs); i++ {
		org, err := database.NewOrganization(c).Create(&orgs[i])
		if err != nil {
			logger.Error(fmt.Sprintf("unable to insert org: %s", err))
			panic(err)
		}
		logger.Info(fmt.Sprintf("inserted org: %v", org))
	}

	memberships, err := getMembershipsFromCsv(users, orgs)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(memberships); i++ {
		_, err = database.NewMembership(c).Create(&memberships[i])
		if err != nil {
			logger.Error(fmt.Sprintf("unable to create membership: %s", err))
			panic(err)
		}
	}
}
