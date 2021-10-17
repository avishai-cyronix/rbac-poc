package rbac

import "github.com/casbin/casbin/v2"

type Loader interface {
	// Enforce a resource act
	Enforce(*Model) (bool, error)
}
type loader struct {
	Enforcer *casbin.Enforcer
}

func NewLoader(modelPath, policyPath string) (Loader, error) {
	e, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, err
	}
	return &loader{
		Enforcer: e,
	}, nil
}

// Enforce a resource act
func (l *loader) Enforce(model *Model) (bool, error) {
	return l.Enforcer.Enforce(model.Role, model.Object, model.Act)
}
