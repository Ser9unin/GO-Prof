package hw09structvalidator

import (
	"strconv"
	"strings"
)

func validateInt(fieldName, fieldTag string, fvInt int64) error {
	var validateErr ValidationErrors

	s := strings.Split(fieldTag, "|")

	for _, v := range s {
		requirement := strings.Split(v, ":")
		if err := checkRequirement(requirement, fieldName); err != nil {
			return err
		}

		switch requirement[0] {
		case "min":
			if minRequirement, err := strconv.Atoi(requirement[1]); err != nil {
				return err
			} else if fvInt < int64(minRequirement) {
				validateErr = append(validateErr, ValidationError{fieldName, ErrMin})
			}
		case "max":
			if maxRequirement, err := strconv.Atoi(requirement[1]); err != nil {
				return err
			} else if fvInt > int64(maxRequirement) {
				validateErr = append(validateErr, ValidationError{fieldName, ErrMax})
			}
		case "in":
			ok, err := inIntSet(requirement[1], fvInt)
			if err != nil {
				validateErr = append(validateErr, ValidationError{fieldName, err})
			}
			if !ok {
				validateErr = append(validateErr, ValidationError{fieldName, ErrIn})
			}
		default:
			validateErr = append(validateErr, ValidationError{fieldName, ErrInvalidRequirement})
		}
	}
	return validateErr
}
