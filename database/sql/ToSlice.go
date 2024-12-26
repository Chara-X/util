package sql

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

func ToSlice[T any](cursor *sql.Rows) []T {
	var rows = []map[string]any{}
	var cols, _ = cursor.Columns()
	for cursor.Next() {
		var row = map[string]any{}
		var dest = make([]any, len(cols))
		for i := range cols {
			dest[i] = new(any)
		}
		cursor.Scan(dest...)
		for i, col := range cols {
			row[col] = *dest[i].(*any)
		}
		rows = append(rows, row)
	}
	// if T is an interface, return []map[string]any.
	if reflect.TypeFor[T]().Kind() == reflect.Interface {
		return any(rows).([]T)
	}
	var objs = make([]T, len(rows))
	// if T is a struct, return []T.
	if reflect.TypeFor[T]().Kind() == reflect.Struct {
		for k, v := range rows {
			var s, _ = json.Marshal(v)
			json.Unmarshal(s, &objs[k])
		}
		return objs
	}
	// if T is a scalar, return []T.
	for i, v := range rows {
		objs[i] = reflect.ValueOf(v[cols[0]]).Convert(reflect.TypeFor[T]()).Interface().(T)
	}
	return objs
}
