//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/enginer/jet/v2/postgres"
)

var FilmActor = newFilmActorTable("dvds", "film_actor", "")

type filmActorTable struct {
	postgres.Table

	//Columns
	ActorID    postgres.ColumnInteger
	FilmID     postgres.ColumnInteger
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FilmActorTable struct {
	filmActorTable

	EXCLUDED filmActorTable
}

// AS creates new FilmActorTable with assigned alias
func (a FilmActorTable) AS(alias string) *FilmActorTable {
	return newFilmActorTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FilmActorTable with assigned schema name
func (a FilmActorTable) FromSchema(schemaName string) *FilmActorTable {
	return newFilmActorTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FilmActorTable with assigned table prefix
func (a FilmActorTable) WithPrefix(prefix string) *FilmActorTable {
	return newFilmActorTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FilmActorTable with assigned table suffix
func (a FilmActorTable) WithSuffix(suffix string) *FilmActorTable {
	return newFilmActorTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFilmActorTable(schemaName, tableName, alias string) *FilmActorTable {
	return &FilmActorTable{
		filmActorTable: newFilmActorTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newFilmActorTableImpl("", "excluded", ""),
	}
}

func newFilmActorTableImpl(schemaName, tableName, alias string) filmActorTable {
	var (
		ActorIDColumn    = postgres.IntegerColumn("actor_id")
		FilmIDColumn     = postgres.IntegerColumn("film_id")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
		allColumns       = postgres.ColumnList{ActorIDColumn, FilmIDColumn, LastUpdateColumn}
		mutableColumns   = postgres.ColumnList{LastUpdateColumn}
	)

	return filmActorTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ActorID:    ActorIDColumn,
		FilmID:     FilmIDColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
