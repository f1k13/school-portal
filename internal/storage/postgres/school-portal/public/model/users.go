//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	ID           uuid.UUID `sql:"primary_key"`
	FirstName    string
	LastName     string
	Email        string
	MiddleName   *string
	RefreshToken string
	Role         string
	PhoneNumber  string
	CreatedAt    *time.Time
}
