// Package models contains the types for schema 'xodb'.
package models

// GENERATED BY XO. DO NOT EDIT.

// MyEnum represents a MySQL enum.
type MyEnum struct {
	TableName  string // table_name
	EnumType   string // enum_type
	EnumValues string // enum_values
}

// MyEnumsBySchema runs a custom query, returning results as MyEnum.
func MyEnumsBySchema(db XODB, schema string) ([]*MyEnum, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`table_name AS table_name, ` +
		`column_name AS enum_type, ` +
		`SUBSTRING(column_type, 6, CHAR_LENGTH(column_type) - 6) AS enum_values ` +
		`FROM information_schema.columns ` +
		`WHERE data_type = 'enum' AND table_schema = ? ` +
		`ORDER BY table_name, column_name`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*MyEnum{}
	for q.Next() {
		me := MyEnum{}

		// scan
		err = q.Scan(&me.TableName, &me.EnumType, &me.EnumValues)
		if err != nil {
			return nil, err
		}

		res = append(res, &me)
	}

	return res, nil
}