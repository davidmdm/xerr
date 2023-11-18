package xerr

import (
	"strings"
)

const defaultIndent = "  "

type MultiErr struct {
	Errors []error
	Label  string
	Indent string
}

func (err MultiErr) Unwrap() []error { return err.Errors }

func (err MultiErr) Error() string {
	if err.Label == "" {
		err.Label = "error"
	}
	switch len(err.Errors) {
	case 0:
		return err.Label
	case 1:
		return err.Label + ": " + err.Errors[0].Error()
	default:
		{
			if err.Indent == "" {
				err.Indent = defaultIndent
			}

			var builder strings.Builder

			builder.WriteString(err.Label + ":")
			for _, e := range err.Errors {
				builder.WriteString("\n" + indent("- "+e.Error(), err.Indent))
			}
			return builder.String()
		}
	}
}

func MultiErrWithIndentFrom(msg, indent string, errs ...error) error {
	var nonNilErrs []error
	for _, err := range errs {
		if err != nil {
			nonNilErrs = append(nonNilErrs, err)
		}
	}
	if len(nonNilErrs) == 0 {
		return nil
	}
	return MultiErr{
		Errors: errs,
		Label:  msg,
		Indent: indent,
	}
}

func MultiErrFrom(msg string, errs ...error) error {
	return MultiErrWithIndentFrom(msg, defaultIndent, errs...)
}

func indent(value, indent string) string {
	return indent + strings.ReplaceAll(value, "\n", "\n"+indent)
}
