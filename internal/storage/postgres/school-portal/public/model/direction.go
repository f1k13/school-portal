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

type Direction struct {
	ID          uuid.UUID `sql:"primary_key"`
	Name        string
	Image       string
	Description string
	CreatedAt   *time.Time
}
