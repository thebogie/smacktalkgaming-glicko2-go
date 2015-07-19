// util
package models

import (
	"crypto/rand"
	//"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
	"io"
	"log"
	"reflect"
	//"mitchgottlieb.com/smacktalkgaming/app/models"
	"errors"
)

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func ConvertToProps(props *neoism.Props, data interface{}) (label string, err error) {

	//what kind of create are we doing?
	log.Println("TYPDOF", reflect.TypeOf(data))
	switch data.(type) {
	case *Game:
		label = "Game"
		log.Println("READ GAME!")
		temp, _ := ToMap(data.(*Game), "")
		*props = temp

	default:
		log.Println("FALL THROUGH")
	}

	return label, err
}

func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	log.Println("V:", v)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		//log.Println("FIELD: + interface", fi, v.Field(i).Interface())
		//if tagv := fi.Tag.Get(tag); tagv != "" {
		// set key of map to value in struct field
		//out[tagv] = v.Field(i).Interface()
		//}

		out[fi.Name] = v.Field(i).Interface()
	}
	return out, nil
}

// convert from map[stirng]interfact{} to a struct
func FillStruct(m map[string]interface{}, s interface{}) error {
	structValue := reflect.ValueOf(s).Elem()

	for name, value := range m {
		structFieldValue := structValue.FieldByName(name)

		if !structFieldValue.IsValid() {
			return fmt.Errorf("No such field: %s in obj", name)
		}

		if !structFieldValue.CanSet() {
			return fmt.Errorf("Cannot set %s field value", name)
		}

		val := reflect.ValueOf(value)
		if structFieldValue.Type() != val.Type() {
			return errors.New("Provided value type didn't match obj field type")
		}

		structFieldValue.Set(val)
	}
	return nil
}
