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
		var elem = map[string]any{}
		var vals = make([]any, len(cols))
		for i := range cols {
			vals[i] = new(any)
		}
		cursor.Scan(vals...)
		for i, col := range cols {
			elem[col] = *vals[i].(*any)
		}
		rows = append(rows, elem)
	}
	if reflect.TypeFor[T]().Kind() == reflect.Interface {
		return any(rows).([]T)
	}
	var objs = make([]T, len(rows))
	if reflect.TypeFor[T]().Kind() == reflect.Struct {
		for k, v := range rows {
			var s, _ = json.Marshal(v)
			json.Unmarshal(s, &objs[k])
		}
		return objs
	}
	for i, v := range rows {
		objs[i] = reflect.ValueOf(v[cols[0]]).Convert(reflect.TypeFor[T]()).Interface().(T)
	}
	return objs
}
