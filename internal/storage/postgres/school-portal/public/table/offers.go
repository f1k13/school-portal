//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Offers = newOffersTable("public", "offers", "")

type offersTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	UserID      postgres.ColumnString
	Price       postgres.ColumnInteger
	DirectionID postgres.ColumnString
	Title       postgres.ColumnString
	Description postgres.ColumnString
	IsOnline    postgres.ColumnBool
	CreatedAt   postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OffersTable struct {
	offersTable

	EXCLUDED offersTable
}

// AS creates new OffersTable with assigned alias
func (a OffersTable) AS(alias string) *OffersTable {
	return newOffersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OffersTable with assigned schema name
func (a OffersTable) FromSchema(schemaName string) *OffersTable {
	return newOffersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OffersTable with assigned table prefix
func (a OffersTable) WithPrefix(prefix string) *OffersTable {
	return newOffersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OffersTable with assigned table suffix
func (a OffersTable) WithSuffix(suffix string) *OffersTable {
	return newOffersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOffersTable(schemaName, tableName, alias string) *OffersTable {
	return &OffersTable{
		offersTable: newOffersTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newOffersTableImpl("", "excluded", ""),
	}
}

func newOffersTableImpl(schemaName, tableName, alias string) offersTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		UserIDColumn      = postgres.StringColumn("user_id")
		PriceColumn       = postgres.IntegerColumn("price")
		DirectionIDColumn = postgres.StringColumn("direction_id")
		TitleColumn       = postgres.StringColumn("title")
		DescriptionColumn = postgres.StringColumn("description")
		IsOnlineColumn    = postgres.BoolColumn("is_online")
		CreatedAtColumn   = postgres.TimestampColumn("created_at")
		allColumns        = postgres.ColumnList{IDColumn, UserIDColumn, PriceColumn, DirectionIDColumn, TitleColumn, DescriptionColumn, IsOnlineColumn, CreatedAtColumn}
		mutableColumns    = postgres.ColumnList{UserIDColumn, PriceColumn, DirectionIDColumn, TitleColumn, DescriptionColumn, IsOnlineColumn, CreatedAtColumn}
	)

	return offersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		UserID:      UserIDColumn,
		Price:       PriceColumn,
		DirectionID: DirectionIDColumn,
		Title:       TitleColumn,
		Description: DescriptionColumn,
		IsOnline:    IsOnlineColumn,
		CreatedAt:   CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
