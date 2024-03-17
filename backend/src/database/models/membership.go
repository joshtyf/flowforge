package models

import "time"

type MembershipModel struct {
	UserId   string    `json:"user_id"`
	OrgId    int       `json:"org_id"`
	JoinedOn time.Time `json:"joined_on"`
}
