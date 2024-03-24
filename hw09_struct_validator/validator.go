package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

var (
	ErrInputNotStruct     = errors.New("input not struct")
	ErrRequirement        = errors.New("requirement not full")
	ErrInvalidRequirement = errors.New("invalid reuqirement")
	ErrLen                = errors.New("length not match requirement")
	ErrRegexp             = errors.New("string not correspond to regexp")
	ErrMin                = errors.New("less than min")
	ErrMax                = errors.New("more than max")
	ErrIn                 = errors.New("not in set")
	ErrEmptyInput         = errors.New("input is empty")
)

func (v ValidationErrors) Error() string {
	var res strings.Builder
	for _, k := range v {
		res.WriteString(fmt.Sprintf("%s: %v ", k.Field, k.Err))
	}
	return res.String()
}

func Validate(v interface{}) error {
	// Place your code here.
	if v == nil {
		return ErrEmptyInput
	}

	objElem := reflect.ValueOf(v).Elem()
	if objElem.Kind() != reflect.Struct {
		return ErrInputNotStruct
	}

	var valErrors ValidationErrors

	objType := objElem.Type()

	for _, field := range reflect.VisibleFields(objType) {

		fieldTag, ok := field.Tag.Lookup("validate")
		if !ok {
			continue
		}

		fieldName := field.Name
		fieldVal := objElem.FieldByName(fieldName)
		fieldType := field.Type

		var err error

		switch fieldType.Kind() {
		case reflect.String:
			fvString := fieldVal.String()
			err = validateString(fieldName, fieldTag, fvString)
		case reflect.Int:
			fvInt := fieldVal.Int()
			err = validateInt(fieldName, fieldTag, fvInt)
		case reflect.Array:
			if fieldType.Elem().String() == "int" {
				for i := 0; i < fieldVal.NumField(); i++ {
					fvInt := fieldVal.Field(i).Int()
					err = validateInt(fieldName, fieldTag, fvInt)
				}
			}
			if fieldType.Elem().String() == "string" {
				for i := 0; i < fieldVal.NumField(); i++ {
					fvString := fieldVal.Field(i).String()
					err = validateString(fieldName, fieldTag, fvString)
				}
			}
		}

		var v ValidationErrors
		if !errors.As(err, &v) {
			return err
		}
		valErrors = append(valErrors, v...)
	}

	return valErrors
}
