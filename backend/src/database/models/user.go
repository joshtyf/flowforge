package models

import (
	"time"
)

type UserModel struct {
	Id             string    `json:"user_id,omitempty"`
	Name           string    `json:"name,omitempty"`
	ConnectionType string    `json:"connection,omitempty"`
	CreatedOn      time.Time `json:"created_on"`
}