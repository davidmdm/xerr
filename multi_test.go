package xerr_test

import (
	"errors"
	"testing"

	"github.com/davidmdm/xerr"
	"github.com/stretchr/testify/require"
)

func TestMultiErrRender(t *testing.T) {
	cases := []struct {
		Name     string
		Err      xerr.MultiErr
		Expected string
	}{
		{
			Name:     "zero value",
			Err:      xerr.MultiErr{},
			Expected: "error",
		},
		{
			Name:     "only message",
			Err:      xerr.MultiErr{Label: "error occured"},
			Expected: "error occured",
		},
		{
			Name: "one error",
			Err: xerr.MultiErr{
				Label:  "error occured",
				Errors: []error{errors.New("one")},
			},
			Expected: "error occured: one",
		},
		{
			Name: "multiple error",
			Err: xerr.MultiErr{
				Label:  "error occured",
				Errors: []error{errors.New("one"), errors.New("two")},
			},
			Expected: "error occured:\n  - one\n  - two",
		},
		{
			Name: "nested",
			Err: xerr.MultiErr{
				Label: "error occured",
				Errors: []error{
					errors.New("one"),
					xerr.MultiErr{
						Label:  "nested error occured",
						Errors: []error{errors.New("a"), errors.New("b")},
					},
				},
			},
			Expected: "error occured:\n  - one\n  - nested error occured:\n    - a\n    - b",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			require.Equal(t, tc.Expected, tc.Err.Error())
		})
	}
}

func TestMultiFrom(t *testing.T) {
	cases := []struct {
		Name     string
		Label    string
		Errs     []error
		Expected string
	}{
		{
			Name:     "nil errors",
			Errs:     []error{nil, nil, nil},
			Expected: "",
		},
		{
			Name:     "msg but no errors",
			Label:    "custom",
			Errs:     []error{nil, nil, nil},
			Expected: "",
		},
		{
			Name:     "msg and single error",
			Label:    "custom",
			Errs:     []error{errors.New("error")},
			Expected: "custom: error",
		},
		{
			Name:     "msg and multiple error",
			Label:    "custom",
			Errs:     []error{errors.New("a"), errors.New("b")},
			Expected: "custom:\n  - a\n  - b",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			err := xerr.MultiErrFrom(tc.Label, tc.Errs...)
			if tc.Expected == "" {
				require.NoError(t, err)
			} else {
				require.Equal(t, tc.Expected, err.Error())
			}
		})
	}
}
