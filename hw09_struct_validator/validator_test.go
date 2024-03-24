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
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	String struct {
		Str string `validate:"in:admin,stuff"`
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
	var valErrors ValidationErrors
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          &User{ID: "123", Name: "Alex", Age: 3, meta: nil},
			expectedErr: ValidationErrors{ValidationError{Field: "Age", Err: ErrMin}},
		},
		{
			in:          &User{ID: "1", Name: "Alex", Age: 18},
			expectedErr: valErrors,
		},
		{
			in:          &User{ID: "123", Email: "mail@google"},
			expectedErr: ValidationErrors{ValidationError{Field: "Email", Err: ErrRegexp}},
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
			in:          &Token{Header: []byte{0, 1, 2}},
			expectedErr: valErrors,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			if err := Validate(tt.in); err != nil {
				fmt.Println(err)
				require.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
