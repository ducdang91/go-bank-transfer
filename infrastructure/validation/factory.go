package validation

import (
	"errors"

	"github.com/ducdang91/go-bank-transfer/adapter/validator"
)

var (
	errInvalidValidatorInstance = errors.New("invalid validator instance")
)

const (
	InstanceGoPlayground int = iota
)

func NewValidatorFactory(instance int) (validator.Validator, error) {
	switch instance {
	case InstanceGoPlayground:
		return NewGoPlayground()
	default:
		return nil, errInvalidValidatorInstance
	}
}
