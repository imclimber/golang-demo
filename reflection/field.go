package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/dgryski/trifles/uuid"
)

// Entity ...
type Entity struct {
	EntityID               string   `gorm:"type:VARCHAR(100);primary_key" tag:"entity_id"`
	EntityType             *int     `gorm:"type:INT;not null" tag:"entity_type"`
	FoundedOn              *string  `gorm:"type:VARCHAR(8)" tag:"founded_on"`
	Sequence               int      `gorm:"type:INT;index;not null" tag:"sequence"`
	RaisedToDateNormalized *float64 `gorm:"type:FLOAT(20,4)" tag:"raised_to_date_normalized"`
}

// GetStructTagAndFieldValue ...
func GetStructTagAndFieldValue() {
	entityID := uuid.UUIDv4()
	entityType := 2
	// foundedOn := "20200101"
	raisedToDateNormalized := 120.12

	var entity = Entity{
		EntityID:   entityID,
		EntityType: &entityType,
		// FoundedOn:  &foundedOn,
		Sequence:               1,
		RaisedToDateNormalized: &raisedToDateNormalized,
	}

	// get type and values of struct
	rType := reflect.TypeOf(entity)
	rValue := reflect.ValueOf(entity)
	fmt.Println("rType: ", rType)
	fmt.Println("rValue: ", rValue)
	fmt.Println("number of struct: ", rType.NumField())

	// get field tag from `TypeOf`return value
	fmt.Println("---------------------- tag ----------------------")
	for i := 0; i < rType.NumField(); i++ {
		tagName := rType.Field(i).Tag.Get("tag")
		fmt.Println(tagName)
	}

	// get field kind from `ValueOf` return value: rValue.Field(i).Kind() (if not pointer)
	// get field kind from `ValueOf` return value: rValue.Field(i).Elem().Kind() (if pointer and not nil)
	fmt.Println("---------------------- name and value ----------------------")
	for i := 0; i < rType.NumField(); i++ {
		// if field is a pointer
		if rValue.Field(i).Kind() == reflect.Ptr {
			// if field pointer not nil
			if !rValue.Field(i).IsNil() {
				fmt.Println("fieldName and kind and value: ", rType.Field(i).Name, rValue.Field(i).Elem().Kind(), rValue.Field(i).Elem().Interface())
			} else {
				fmt.Println("nil value of field: ", rType.Field(i).Name)

				// 测试，如果 rValue.Field(i) 为空，依然可以调用 Elem() 方法，但是调用  Elem() 后不能继续调用其他方法
				fmt.Println("fieldName and kind and value: ", rType.Field(i).Name, rValue.Field(i).Elem())
				fmt.Println("fieldName and kind and value: ", rType.Field(i).Name, rValue.Field(i).Elem().FieldByName("ID").String())
			}

			continue
		}

		fmt.Println("fieldName and kind and value: ", rType.Field(i).Name, rValue.Field(i).Kind(), rValue.Field(i).Interface())
	}

	fmt.Println("---------------------- sql ----------------------")
	var buffer bytes.Buffer
	buffer.WriteString("insert into entity_collection_companies(entity_id, entity_type, founded_on, sequence, raised_to_date_normalized) values")
	for i := 0; i < rType.NumField(); i++ {
		operator := "%v, "

		// if field is a pointer
		if rValue.Field(i).Kind() == reflect.Ptr {
			// if field pointer not nil
			if !rValue.Field(i).IsNil() {
				if rValue.Field(i).Elem().Kind().String() == "string" {
					operator = "'%v', "
				}

				buffer.WriteString(fmt.Sprintf(operator, rValue.Field(i).Elem().Interface()))
			} else {
				buffer.WriteString(fmt.Sprintf(operator, "NULL"))
			}

			continue
		}

		if rValue.Field(i).Kind().String() == "string" {
			operator = "'%v', "
		}
		buffer.WriteString(fmt.Sprintf(operator, rValue.Field(i).Interface()))
	}

	sqlString := buffer.String()
	sqlString = strings.TrimRight(sqlString, ", ")
	fmt.Println("sql: ", "("+sqlString+")")
}
