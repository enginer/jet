package metadata

import (
	"database/sql"
	"fmt"
)

// SchemaMetaData struct
type SchemaMetaData struct {
	TableInfos []MetaData
	EnumInfos  []MetaData
}

// GetSchemaInfo returns schema information from db connection.
func GetSchemaInfo(db *sql.DB, schemaName string, querySet DialectQuerySet) (schemaInfo SchemaMetaData, err error) {

	schemaInfo.TableInfos, err = getTableInfos(db, querySet, schemaName)

	if err != nil {
		return
	}

	schemaInfo.EnumInfos, err = querySet.GetEnumsMetaData(db, schemaName)

	if err != nil {
		return
	}

	fmt.Println("	FOUND", len(schemaInfo.TableInfos), "table(s), ", len(schemaInfo.EnumInfos), "enum(s)")

	return
}

func getTableInfos(db *sql.DB, querySet DialectQuerySet, schemaName string) ([]MetaData, error) {

	rows, err := db.Query(querySet.ListOfTablesQuery(), schemaName)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := []MetaData{}
	for rows.Next() {
		var tableName string

		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		tableInfo, err := GetTableInfo(db, querySet, schemaName, tableName)

		if err != nil {
			return nil, err
		}

		ret = append(ret, tableInfo)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return ret, nil
}