package hw09structvalidator

import (
	"fmt"
	"strconv"
	"strings"
)

func checkRequirement(requirement []string, fieldName string) error {
	if len(requirement) != 2 {
		return fmt.Errorf("field %s: %w", fieldName, ErrRequirement)
	}
	if requirement[1] == "" {
		return fmt.Errorf("field %s: %w", fieldName, ErrRequirement)
	}
	return nil
}

func inStringSet(requirement, fvString string) bool {
	set := strings.Split(requirement, ",")
	for _, v := range set {
		if fvString == v {
			return true
		}
	}
	return false
}

func inIntSet(requirement string, fvInt int64) (bool, error) {
	set := strings.Split(requirement, ",")
	for _, v := range set {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return false, ErrInvalidRequirement
		}
		if fvInt == int64(intV) {
			return true, nil
		}
	}
	return false, nil
}
