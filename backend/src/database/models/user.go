package models

import (
	"time"
)

type UserModel struct {
	UserId           string               `json:"user_id"`
	Name             string               `json:"name"`
	IdentityProvider string               `json:"identity_provider"`
	Organisations    []*OrganisationModel `json:"organisations,omitempty"`
	CreatedOn        time.Time            `json:"created_on"`
}
