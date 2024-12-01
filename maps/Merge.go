package maps

import "reflect"

func Merge(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		if _, ok := m1[k]; ok && reflect.TypeOf(m1[k]).Kind() == reflect.Map {
			m1[k] = Merge(m1[k].(map[string]interface{}), v.(map[string]interface{}))
		} else {
			m1[k] = v
		}
	}
	return m1
}
