package errors

import "reflect"

// Error implement error interface to create
// custom error type
type Error struct {
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Field       string      `json:"field,omitempty"`
	Instance    interface{} `json:"instance,omitempty"`
}

func (err Error) Error() string {
	return err.Description
}

func New(code, description, field string, instance interface{}) Error {
	return Error{
		Code:        code,
		Description: description,
		Field:       field,
		Instance:    instance,
	}
}

// CastError cast err to type Error if err is
// indeed are from Error type
func CastError(err error) (bool, Error) {
	val := reflect.ValueOf(err)
	if val.IsNil() {
		return false, Error{}
	}

	elem := val.Elem()
	for elem.Kind() == reflect.Pointer {
		elem = elem.Elem()
	}

	elemT := elem.Type()

	errT := reflect.TypeOf(Error{})
	if elemT.String() == errT.String() {
		return true, elem.Interface().(Error)
	}

	return false, Error{}
}
