package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:3"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	String struct {
		Str string `validate:"in:admin,stuff"`
	}

	IntSlice struct {
		IntS []int `validate:"min:18|max:50"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          &User{ID: "123", Name: "Alex", Age: 3, meta: nil},
			expectedErr: ValidationErrors{ValidationError{Field: "Age", Err: ErrMin}},
		},
		{
			in: &User{
				ID:     "123",
				Name:   "Alex",
				Age:    3,
				Email:  "user@domen.com",
				Role:   "admin",
				Phones: []string{"7999123456", "+7888123456"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Age", Err: ErrMin},
				ValidationError{Field: "Phones", Err: ErrLen},
			},
		},
		{
			in:          &User{ID: "1", Name: "Alex", Age: 18},
			expectedErr: ValidationErrors{ValidationError{Field: "ID", Err: ErrLen}},
		},
		{
			in: &User{ID: "123", Email: "mail@google"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Age", Err: ErrMin},
				ValidationError{Field: "Email", Err: ErrRegexp},
			},
		},
		{
			in:          &App{Version: "1.2.3.456"},
			expectedErr: ValidationErrors{ValidationError{Field: "Version", Err: ErrLen}},
		},
		{
			in:          &Response{Code: 307},
			expectedErr: ValidationErrors{ValidationError{Field: "Code", Err: ErrIn}},
		},
		{
			in:          &String{Str: "user"},
			expectedErr: ValidationErrors{ValidationError{Field: "Str", Err: ErrIn}},
		},
		{
			in: &IntSlice{IntS: []int{1, 18, 20, 50, 55}},
			expectedErr: ValidationErrors{
				ValidationError{Field: "IntS", Err: ErrMin},
				ValidationError{Field: "IntS", Err: ErrMax},
			},
		},
	}

	for i, testcase := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			testcase := testcase
			// t.Parallel()
			err := Validate(testcase.in)
			fmt.Println(testcase.expectedErr, " ", err)
			require.Equal(t, testcase.expectedErr, err)
		})
	}
}
