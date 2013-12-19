package main

import (
	"errors"
	"reflect"
)

type Secondary struct {
	String string  `request:"string_value"`
	Int    float64 `request:"int_value"`
	Bool   bool    `request:"bool_value"`
}

type Session struct {
	String string    `request:"string_value"`
	Int    float64   `request:"int_value"`
	Bool   bool      `request:"bool_value"`
	Obj    Secondary `request:"obj_value"`
}

func ValidateInput(d map[string]interface{}, s interface{}) error {
	sVal := reflect.ValueOf(s).Elem()
	for i := 0; i < sVal.NumField(); i++ {
		field := reflect.TypeOf(s).Elem().Field(i)
		tag := field.Tag.Get("request")
		if val, ok := d[tag]; ok {
			f := sVal.Field(i)
			if f.IsValid() && f.CanSet() {
				if reflect.TypeOf(val).Kind() == reflect.Map {
					item := reflect.New(f.Type()).Elem().Addr().Interface()
					err := ValidateInput(val.(map[string]interface{}), item)
					if err != nil {
						return err
					}
					f.Set(reflect.ValueOf(item).Elem())
				} else {
					if f.Type() == reflect.TypeOf(val) {
						switch val.(type) {
						case float64:
							f.SetFloat(val.(float64))
						case string:
							f.SetString(val.(string))
						case bool:
							f.SetBool(val.(bool))
						}
					} else {
						return errors.New("input data type does not match the struct feild")
					}
				}
			} else {
				return errors.New("cannot set values; struct must have exported fields")
			}
		} else {
			return errors.New("required data is missing")
		}
	}
	return nil
}
