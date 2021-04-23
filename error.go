package gassert

import (
	"errors"
	"fmt"
)

var (
	numLessError            = errors.New("num less error")
	numLessOrEqualsError    = errors.New("num less or equals error")
	numGreaterError         = errors.New("num greater error")
	numGreaterOrEqualsError = errors.New("num greater or equals error")
)

type Errors []error

func (errs Errors) Error() string {
	msg := fmt.Sprintln("gAssertError")
	msg += fmt.Sprintln("Details")
	for i, err := range errs {
		msg += fmt.Sprintf("[%d] %s\n", i, err.Error())
	}
	return msg
}

type zeroValueError string

func (e zeroValueError) Error() string {
	return fmt.Sprintf("zero-value error in %s", string(e))
}

type equalsError string

func (e equalsError) Error() string {
	return fmt.Sprintf("equals error in %s", string(e))
}

type notEqualsError string

func (e notEqualsError) Error() string {
	return fmt.Sprintf("not equals error in %s", string(e))
}

type lenEqualsError string

func (e lenEqualsError) Error() string {
	return fmt.Sprintf("len equals error in %s", string(e))
}

type lenNotEqualsError string

func (e lenNotEqualsError) Error() string {
	return fmt.Sprintf("len not equals error in %s", string(e))
}

type lenLessError string

func (e lenLessError) Error() string {
	return fmt.Sprintf("len less error in %s", string(e))
}

type lenLessOrEqualsError string

func (e lenLessOrEqualsError) Error() string {
	return fmt.Sprintf("len less or equals error in %s", string(e))
}

type lenGreaterError string

func (e lenGreaterError) Error() string {
	return fmt.Sprintf("len greater equals error in %s", string(e))
}

type lenGreaterOrEqualsError string

func (e lenGreaterOrEqualsError) Error() string {
	return fmt.Sprintf("len greater or equals error in %s", string(e))
}

type capEqualsError string

func (e capEqualsError) Error() string {
	return fmt.Sprintf("cap equals error in %s", string(e))
}

type capNotEqualsError string

func (e capNotEqualsError) Error() string {
	return fmt.Sprintf("cap not equals error in %s", string(e))
}

type capLessError string

func (e capLessError) Error() string {
	return fmt.Sprintf("cap less error in %s", string(e))
}

type capLessOrEqualsError string

func (e capLessOrEqualsError) Error() string {
	return fmt.Sprintf("cap less or equals error in %s", string(e))
}

type capGreaterError string

func (e capGreaterError) Error() string {
	return fmt.Sprintf("cap greater error in %s", string(e))
}

type capGreaterOrEqualsError string

func (e capGreaterOrEqualsError) Error() string {
	return fmt.Sprintf("cap greater or equals error in %s", string(e))
}
