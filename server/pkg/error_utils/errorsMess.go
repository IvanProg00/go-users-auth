package error_utils

import (
	"github.com/pkg/errors"
)

func IsEmptyError(field string) error {
	return errors.Errorf("%s is empty", field)
}
