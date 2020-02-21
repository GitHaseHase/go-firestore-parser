package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// Test for firestore value parser
func Test() {
	bytes, err := ioutil.ReadFile("example.json")
	if err != nil {
		panic(err)
	}

	var inter map[string]interface{}
	if err := json.Unmarshal(bytes, &inter); err != nil {
		panic(err)
	}
	fields := inter["fields"]
	// ParseFirestoreValue(&fields)
	// result := ParseFirestoreValue(fields)
	result := ParseFirestoreValue(&fields)
	fmt.Println(reflect.Indirect(reflect.ValueOf(result)).Interface())
}
