package models

import (
	"errors"
	"time"
)

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

func GetRoleFromString(roleStr string) (Role, error) {
	roleMap := map[string]Role{
		"Owner":  Owner,
		"Admin":  Admin,
		"Member": Member,
	}

	role, found := roleMap[roleStr]
	if !found {
		return "", errors.New("invalid role")
	}

	return role, nil
}
