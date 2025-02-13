package tests

import (
	"reflect"

	"github.com/go-faker/faker/v4"
)

var fkUUID = faker.UUID{}
var uuid1 = getFakeUUID()
var uuid2 = getFakeUUID()
var post1 = getFakeUUID()
var post2 = getFakeUUID()

func getFakeUUID() string {
	uuidStr, err := fkUUID.Hyphenated(reflect.Value{})
	if err != nil {
		// this is a very unlikely case
		return "dd81db0e-4530-4147-9951-7e9c7998ac34"
	} else {
		return uuidStr.(string)
	}
}
