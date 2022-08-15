package utils

import "reflect"

// Get object name from reflect
func GetObjectName(object interface{}) string {
	// Because struct and interface get name is different, so we first try get name, if nil, then get elem name
	objectName := reflect.TypeOf(object).Name()
	if objectName == "" {
		objectName = reflect.TypeOf(object).Elem().Name()
	}
	return objectName
}
