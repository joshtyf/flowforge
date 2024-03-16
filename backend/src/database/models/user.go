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

type OrganisationModel struct {
	Id        int       `json:"org_id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	CreatedOn time.Time `json:"created_on"`
}

type MembershipModel struct {
	UserId   string    `json:"user_id"`
	OrgId    int       `json:"org_id"`
	JoinedOn time.Time `json:"joined_on"`
}
