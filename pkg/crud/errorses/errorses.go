package errorses

import "fmt"

type QueryError struct {
	Query string
	Err   error
}

type DbError struct {
	Err error
}

type ModelError struct {
	Value string
	Err error
}

func (receiver *QueryError) Unwrap() error {
	return receiver.Err
}

func (receiver *QueryError) Error() string {
	return fmt.Sprintf("can't execute query %s: %s", receiver.Query, receiver.Err.Error())
}

func (receiver *DbError) Error() string {
	return fmt.Sprintf("can't handle db operation: %v", receiver.Err.Error())
}

func (receiver *DbError) Unwrap() error {
	return receiver.Err
}

func (receiver *ModelError) Unwrap() error {
	return receiver.Err
}

func (receiver *ModelError) Error() string {
	return fmt.Sprintf("can't execute query %s: %s", receiver.Value, receiver.Err.Error())
}

func QueryErrors(query string, err error) *QueryError {
	return &QueryError{Query: query, Err: err}
}

func ModelErrors(value string, err error) *ModelError {
	return &ModelError{Value: value, Err: err}
}

func DbErrors(err error) *DbError {
	return &DbError{Err: err}
}