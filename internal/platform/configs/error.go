package configs

import "github.com/kumarabd/gokit/errors"

func ErrLocal(err error) error {
	return errors.New("CONFIG.INIT.ERROR", "Local, error: "+err.Error())
}
