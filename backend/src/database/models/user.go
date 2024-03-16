package models

import (
	"time"
)

type UserModel struct {
	Id             string               `json:"user_id"`
	Name           string               `json:"name"`
	ConnectionType string               `json:"connection"`
	Organisations  []*OrganisationModel `json:"organisations,omitempty"`
	CreatedOn      time.Time            `json:"created_on"`
}
