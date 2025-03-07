package user

import (
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

var UserTable = postgres.NewTable("user_schema", "user_table", "user")

type User struct {
	id    uuid.UUID `sql: "primary_key"`
	name  string
	email string
}
