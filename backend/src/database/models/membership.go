package models

import "time"

type Role string

const (
	Owner  Role = "Owner"
	Admin  Role = "Admin"
	Member Role = "Member"
)

type MembershipModel struct {
	UserId   string    `json:"user_id"`
	OrgId    int       `json:"org_id"`
	Role     Role      `json:"role"`
	JoinedOn time.Time `json:"joined_on"`
	Deleted  bool      `json:"deleted"`
}
