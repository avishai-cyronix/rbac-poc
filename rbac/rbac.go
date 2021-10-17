package rbac

import (
	"errors"
)

var (
	NotPermittedErr = errors.New("not permitted")
)

type enforcer struct {
	enforcer Loader
}

func NewEnforcer(loader Loader) Enforcer {
	return &enforcer{
		enforcer: loader,
	}
}

//Allowed checks if subject can act upon a resource
func (e *enforcer) Allowed(model *Model) error {
	ok, err := e.enforcer.Enforce(model)
	if err != nil {
		return err
	}

	if !ok {
		return NotPermittedErr
	}

	return nil
}
