package config

import (
	"errors"
	"reflect"
	"strconv"
)

func Unmarshal(file_path string, object interface{}) error {
	_config, err := NewConfig(file_path)
	if err != nil {
		return err
	}

	myref := reflect.ValueOf(object).Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		value := _config.GetString(typeOfType.Field(i).Name)
		if value == "" {
			return errors.New("Can not find config " + typeOfType.Field(i).Name)
		}
		switch field.Type().Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.CanSet() {
				i, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				field.SetInt(int64(i))
			}
			break
		case reflect.String:
			if field.CanSet() {
				field.SetString(value)
			}
			break
		default:
			return errors.New("Unknow type " + field.Type().String() + " for " + typeOfType.Field(i).Name)
		}
	}
	return nil
}
