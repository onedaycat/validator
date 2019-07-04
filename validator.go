package validator

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"github.com/onedaycat/errors"
)

var (
	emptyStr          = ""
	emailPatern       = regexp.MustCompile("^[A-Z0-9._%-]+@[A-Z0-9.-]+\\.[A-Z]{2,4}$")
	dateiso8601Patern = regexp.MustCompile("^(\\d{4})-(\\d{2})-(\\d{2})T(\\d{2}):(\\d{2}):(\\d{2})(Z|([+\\-])\\d{2}(:?\\d{2})?)$")
)

type ValidateError struct {
	Name string
	Err  error
}

type Pointer interface{}

//go:generate mockery -name Validator
type Validator interface {
	IsValid() bool
	HasError() bool
	GetError() error
	Wrap(errors.Error) errors.Error
	SetError(msg string)

	Required(name string, val Pointer, msg ...interface{})
	NotEmptyString(name string, val string, msg ...interface{})
	NotEmptyInt(name string, val int, msg ...interface{})
	NotEmptyInt64(name string, val int64, msg ...interface{})
	NotEmptyFloat64(name string, val float64, msg ...interface{})
	NotEmptyBool(name string, val bool, msg ...interface{})
	Confirm(name string, val interface{}, confirmName string, confirmValue interface{}, msg ...interface{})
	MaxInt(name string, val int, max int, msg ...interface{})
	MaxInt64(name string, val int64, max int64, msg ...interface{})
	MaxFloat64(name string, val float64, max float64, msg ...interface{})
	MaxString(name string, val string, max int, msg ...interface{})
	MinInt(name string, val int, min int, msg ...interface{})
	MinInt64(name string, val int64, min int64, msg ...interface{})
	MinFloat64(name string, val float64, min float64, msg ...interface{})
	MinString(name string, val string, min int, msg ...interface{})
	EqualString(name string, val string, size int, msg ...interface{})
	EqualInt(name string, val int, equal int, msg ...interface{})
	RangeInt(name string, val int, min, max int, msg ...interface{})
	RangeString(name string, val string, min, max int, msg ...interface{})
	RangeInt64(name string, val int64, min, max int64, msg ...interface{})
	RangeFloat64(name string, val float64, min, max float64, msg ...interface{})
	InString(name string, val string, list []string, msg ...interface{})
	InInt(name string, val int, list []int, msg ...interface{})
	InInt64(name string, val int64, list []int64, msg ...interface{})
	InFloat64(name string, val float64, list []float64, msg ...interface{})
	Email(name string, val string, msg ...interface{})
	ISO8601DataTime(name string, val string, msg ...interface{})
}

type validator struct {
	isError    bool
	errMessage string
}

func New() Validator {
	return &validator{}
}

func (v *validator) IsValid() bool {
	return !v.isError
}

func (v *validator) HasError() bool {
	return v.isError
}

func (v *validator) GetError() error {
	return errors.New(v.errMessage)
}

func (v *validator) Wrap(err errors.Error) errors.Error {
	if v.isError {
		return err.WithMessage(v.errMessage)
	}

	return nil
}

func (v *validator) SetError(msg string) {
	v.isError = true
	v.errMessage = msg
}

func (v *validator) setErr(msg string, customMsg []interface{}) {
	v.isError = true
	if len(customMsg) > 0 {
		v.errMessage = fmt.Sprint(customMsg...)
	} else {
		v.errMessage = msg
	}
}

func (v *validator) Required(name string, val Pointer, msg ...interface{}) {
	if val == nil {
		v.setErr(fmt.Sprintf("%s is required", name), msg)
	}
}

func (v *validator) NotEmptyString(name string, val string, msg ...interface{}) {
	if val == emptyStr {
		v.setErr(fmt.Sprintf("%s must not empty", name), msg)
	}
}

func (v *validator) NotEmptyInt(name string, val int, msg ...interface{}) {
	if val == 0 {
		v.setErr(fmt.Sprintf("%s must not 0", name), msg)
	}
}

func (v *validator) NotEmptyInt64(name string, val int64, msg ...interface{}) {
	if val == 0 {
		v.setErr(fmt.Sprintf("%s must not 0", name), msg)
	}
}

func (v *validator) NotEmptyFloat64(name string, val float64, msg ...interface{}) {
	if val == 0 {
		v.setErr(fmt.Sprintf("%s must not 0", name), msg)
	}
}

func (v *validator) NotEmptyBool(name string, val bool, msg ...interface{}) {
	if !val {
		v.setErr(fmt.Sprintf("%s must not be false", name), msg)
	}
}

func (v *validator) Confirm(name string, val interface{}, confirmName string, confirmValue interface{}, msg ...interface{}) {
	if val != confirmValue {
		v.setErr(fmt.Sprintf("%s value must equal to %s", name, confirmName), msg)
	}
}

func (v *validator) MaxInt(name string, val int, max int, msg ...interface{}) {
	if val > max {
		v.setErr(fmt.Sprintf("%s must not more than %d", name, max), msg)
	}
}

func (v *validator) MaxInt64(name string, val int64, max int64, msg ...interface{}) {
	if val > max {
		v.setErr(fmt.Sprintf("%s must not more than %d", name, max), msg)
	}
}

func (v *validator) MaxFloat64(name string, val float64, max float64, msg ...interface{}) {
	if val > max {
		v.setErr(fmt.Sprintf("%s must not more than %v", name, max), msg)
	}
}

func (v *validator) MaxString(name string, val string, max int, msg ...interface{}) {
	if utf8.RuneCountInString(val) > max {
		v.setErr(fmt.Sprintf("%s must not more than %d characters", name, max), msg)
	}
}

func (v *validator) MinInt(name string, val int, min int, msg ...interface{}) {
	if val > min {
		v.setErr(fmt.Sprintf("%s must not less than %d", name, min), msg)
	}
}

func (v *validator) MinInt64(name string, val int64, min int64, msg ...interface{}) {
	if val > min {
		v.setErr(fmt.Sprintf("%s must not less than %d", name, min), msg)
	}
}

func (v *validator) MinFloat64(name string, val float64, min float64, msg ...interface{}) {
	if val > min {
		v.setErr(fmt.Sprintf("%s must not less than %v", name, min), msg)
	}
}

func (v *validator) MinString(name string, val string, min int, msg ...interface{}) {
	if utf8.RuneCountInString(val) < min {
		v.setErr(fmt.Sprintf("%s must not more than %d characters", name, min), msg)
	}
}

func (v *validator) EqualString(name string, val string, size int, msg ...interface{}) {
	if utf8.RuneCountInString(val) != size {
		v.setErr(fmt.Sprintf("%s must equal %d characters", name, size), msg)
	}
}

func (v *validator) EqualInt(name string, val int, equal int, msg ...interface{}) {
	if val != equal {
		v.setErr(fmt.Sprintf("%s must equal %d", name, equal), msg)
	}
}

func (v *validator) RangeInt(name string, val int, min, max int, msg ...interface{}) {
	if val >= min && val <= max {
		return
	}

	v.setErr(fmt.Sprintf("%s must in range %d and %d", name, min, max), msg)
}

func (v *validator) RangeString(name string, val string, min, max int, msg ...interface{}) {
	steLen := utf8.RuneCountInString(val)
	if steLen >= min && steLen <= max {
		return
	}

	v.setErr(fmt.Sprintf("%s must in range %d and %d characters", name, min, max), msg)
}

func (v *validator) RangeInt64(name string, val int64, min, max int64, msg ...interface{}) {
	if val >= min && val <= max {
		return
	}

	v.setErr(fmt.Sprintf("%s must in range %d and %d", name, min, max), msg)
}

func (v *validator) RangeFloat64(name string, val float64, min, max float64, msg ...interface{}) {
	if val >= min && val <= max {
		return
	}

	v.setErr(fmt.Sprintf("%s must in range %v and %v", name, min, max), msg)
}

func (v *validator) InString(name string, val string, list []string, msg ...interface{}) {
	for _, k := range list {
		if k == val {
			return
		}
	}

	v.setErr(fmt.Sprintf("%s must be in %v", name, list), msg)
}

func (v *validator) InInt(name string, val int, list []int, msg ...interface{}) {
	for _, k := range list {
		if k == val {
			return
		}
	}

	v.setErr(fmt.Sprintf("%s must be in %v", name, list), msg)
}

func (v *validator) InInt64(name string, val int64, list []int64, msg ...interface{}) {
	for _, k := range list {
		if k == val {
			return
		}
	}

	v.setErr(fmt.Sprintf("%s must be in %v", name, list), msg)
}

func (v *validator) InFloat64(name string, val float64, list []float64, msg ...interface{}) {
	for _, k := range list {
		if k == val {
			return
		}
	}

	v.setErr(fmt.Sprintf("%s must be in %v", name, list), msg)
}

func (v *validator) Email(val string, name string, msg ...interface{}) {
	if val == emptyStr {
		v.setErr(fmt.Sprintf("%s must be email format", name), msg)
		return
	}

	if !emailPatern.MatchString(val) {
		v.setErr(fmt.Sprintf("%s must be email format", name), msg)
	}
}

func (v *validator) ISO8601DataTime(val string, name string, msg ...interface{}) {
	if val == emptyStr {
		v.setErr(fmt.Sprintf("%s must be iso8601 date time format", name), msg)
	}

	if !dateiso8601Patern.MatchString(val) {
		v.setErr(fmt.Sprintf("%s must be iso8601 date time format", name), msg)
	}
}
