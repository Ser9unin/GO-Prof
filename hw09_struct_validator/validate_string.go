package hw09structvalidator

import (
	"regexp"
	"strconv"
	"strings"
)

func validateString(fieldName, fieldTag, fvString string) error {
	var validateErr ValidationErrors

	if fvString == "" {
		return nil
	}

	s := strings.Split(fieldTag, "|")

	for _, v := range s {
		requirement := strings.Split(v, ":")
		if err := checkRequirement(requirement, fieldName); err != nil {
			return err
		}

		switch requirement[0] {
		case "len":
			if lenRequirement, err := strconv.Atoi(requirement[1]); err != nil {
				return err
			} else if len(fvString) != lenRequirement {
				validateErr = append(validateErr, ValidationError{fieldName, ErrLen})
			}
		case "regexp":
			if reg, err := regexp.Compile(requirement[1]); err != nil {
				return err
			} else if !reg.Match([]byte(fvString)) {
				validateErr = append(validateErr, ValidationError{fieldName, ErrRegexp})
			}
		case "in":
			if ok := inStringSet(requirement[1], fvString); !ok {
				validateErr = append(validateErr, ValidationError{fieldName, ErrIn})
			}
		default:
			validateErr = append(validateErr, ValidationError{fieldName, ErrInvalidRequirement})
		}
	}
	return validateErr
}
