package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/joshtyf/flowforge/src/database"
	"github.com/joshtyf/flowforge/src/database/client"
	"github.com/joshtyf/flowforge/src/database/models"
	"github.com/joshtyf/flowforge/src/helper"
	"github.com/joshtyf/flowforge/src/logger"
)

func main() {
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
		helper.CreateUsersInAuth0(users, passwords)
	} else {
		helper.GetUserIdForUsers(users)
	}

	for i := 0; i < len(users); i++ {
		user, err := database.NewUser(c).Create(&users[i])
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

func getUsersFromCsv() ([]models.UserModel, []string, error) {
	file, err := os.Open("./database/postgres_seed/" + os.Getenv("USER_SEED_FILENAME"))
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)

	var users []models.UserModel
	var passwords []string

	// skip first row
	_, err = reader.Read()
	if err != nil {
		return nil, nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, nil, err
		} else if err != nil {
			break
		}

		um := models.UserModel{
			Email:            record[0],
			Name:             record[1],
			IdentityProvider: record[2],
		}

		users = append(users, um)
		passwords = append(passwords, record[3])
	}
	return users, passwords, nil
}

func getOrgsFromCsv(users []models.UserModel) ([]models.OrganizationModel, error) {
	file, err := os.Open("./database/postgres_seed/" + os.Getenv("ORG_SEED_FILENAME"))
	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)

	var orgs []models.OrganizationModel

	// skip first row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		} else if err != nil {
			break
		}

		var ownerId string
		for i := 0; i < len(users); i++ {
			if record[1] == users[i].Email {
				ownerId = users[i].UserId
				break
			}
		}

		om := models.OrganizationModel{
			Name:  record[0],
			Owner: ownerId,
		}

		orgs = append(orgs, om)
	}

	return orgs, nil
}

func getMembershipsFromCsv(users []models.UserModel, orgs []models.OrganizationModel) ([]models.MembershipModel, error) {
	file, err := os.Open("./database/postgres_seed/" + os.Getenv("MEMBERSHIP_SEED_FILENAME"))
	if err != nil {
		return nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)

	var memberships []models.MembershipModel

	// skip first row
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		} else if err != nil {
			break
		}

		var userId string
		var orgId int
		for i := 0; i < len(users); i++ {
			if record[0] == users[i].Email {
				userId = users[i].UserId
				break
			}
		}

		for i := 0; i < len(orgs); i++ {
			if record[1] == orgs[i].Name {
				orgId = orgs[i].OrgId
			}
		}

		role, err := models.GetRoleFromString(record[2])
		if err != nil {
			return nil, err
		}
		mm := models.MembershipModel{
			UserId: userId,
			OrgId:  orgId,
			Role:   role,
		}

		memberships = append(memberships, mm)
	}

	return memberships, nil
}
