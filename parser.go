package parser

import (
	"reflect"
	"strconv"
)

// ParseFirestoreValue is parser for the Firestore REST API JSON
func ParseFirestoreValue(value interface{}) interface{} {
	var (
		prop      = GetFirestoreProp(value)
		propName  string
		propValue interface{}
	)
	if prop != nil {
		propName = *prop
		propValue = value.(map[string]interface{})[propName]
	}
	if propName == "integerValue" {
		i, _ := strconv.Atoi(propValue.(string))
		return i
	} else if propName == "geoPointValue" {
		geoPoint := map[string]float64{"latitude": 0, "longitude": 0}
		for _, v := range reflect.ValueOf(propValue).MapKeys() {
			key := v.String()
			geoPoint[key] = propValue.(map[string]interface{})[key].(float64)
		}
		return geoPoint
	} else if propName == "arrayValue" {
		array := propValue.(map[string]interface{})["values"]
		ary := []interface{}{}
		if reflect.ValueOf(array).Kind() == reflect.Slice {
			for _, val := range array.([]interface{}) {
				ary = append(ary, ParseFirestoreValue(val))
			}
		}
		return ary
	} else if propName == "mapValue" {
		obj := propValue.(map[string]interface{})["fields"]
		m := map[string]interface{}{}
		if reflect.ValueOf(obj).Kind() == reflect.Map {
			for key, value := range obj.(map[string]interface{}) {
				m[key] = ParseFirestoreValue(value)
			}
		}
		return m
	} else if prop != nil {
		return propValue
	} else if reflect.ValueOf(value).Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, v := range reflect.ValueOf(value).MapKeys() {
			key := v.String()
			m[key] = ParseFirestoreValue(value.(map[string]interface{})[v.String()])
		}
		return m
	} else if reflect.ValueOf(value).Kind() == reflect.Array {
		ary := []interface{}{}
		for _, value := range value.([]interface{}) {
			ary = append(ary, value)
		}
		return ary
	} else if reflect.ValueOf(value).Kind() == reflect.Ptr {
		return ParseFirestoreValue(reflect.Indirect(reflect.ValueOf(value)).Interface())
	}
	// right back if non Firestore REST API JSON
	return value
}

// GetFirestoreProp is getting the Firestore REST API JSON property names
func GetFirestoreProp(value interface{}) (prop *string) {
	var val map[string]interface{}
	fieldNames := []string{
		"booleanValue",
		"stringValue",
		"arrayValue",
		"mapValue",
		"doubleValue",
		"integerValue",
		"geoPointValue",
		"timestampValue",
		"referenceValue",
		"nullValue",
	}
	if reflect.ValueOf(value).Kind() == reflect.Map {
		val = value.(map[string]interface{})
	}
	for key := range val {
		for _, field := range fieldNames {
			if field == key {
				_field := field
				return &_field
			}
		}
	}
	return prop
}
