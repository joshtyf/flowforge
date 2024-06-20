package seeders

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/joshtyf/flowforge/src/database/models"
)

const DATA_DIR = "./database/seed/data"

func getUsersFromCsv() ([]*models.UserModel, []string, error) {
	file, err := os.Open(DATA_DIR + "/" + os.Getenv("USER_SEED_FILENAME"))
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	reader := csv.NewReader(file)

	var users []*models.UserModel
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

		um := &models.UserModel{
			UserId:           record[0],
			Email:            record[1],
			Name:             record[2],
			IdentityProvider: record[3],
		}

		users = append(users, um)
		passwords = append(passwords, record[4])
	}
	return users, passwords, nil
}

func getOrgsFromCsv(users []*models.UserModel) ([]models.OrganizationModel, error) {
	file, err := os.Open(DATA_DIR + "/" + os.Getenv("ORG_SEED_FILENAME"))
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

func getMembershipsFromCsv(users []*models.UserModel, orgs []models.OrganizationModel) ([]models.MembershipModel, error) {
	file, err := os.Open(DATA_DIR + "/" + os.Getenv("MEMBERSHIP_SEED_FILENAME"))
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
