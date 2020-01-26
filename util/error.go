package util

import (
	"errors"
)

func ErrorNew(msg string) error {
	return errors.New(msg)
}

func ErrorWrap(err error, msg string) error {
	return errors.New(msg + "," + err.Error())
}
