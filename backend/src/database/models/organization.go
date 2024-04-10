package models

import "time"

type OrganizationModel struct {
	OrgId     int       `json:"org_id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	CreatedOn time.Time `json:"created_on"`
	Deleted   bool      `json:"deleted"`
}
