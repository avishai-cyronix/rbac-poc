package rbac_test

import (
	"github.com/stretchr/testify/assert"
	"my-module/rbac"
	"testing"
)

type MockLoader struct {
	Allowed bool
}

func (l *MockLoader) Enforce(model *rbac.Model) (bool, error) {
	return l.Allowed, nil
}

//TODO: Add integration tests with the actual policy definition

func TestEnforcer_Allowed(t *testing.T) {
	loader := &MockLoader{
		Allowed: true,
	}
	enforcer := rbac.NewEnforcer(loader)
	model := &rbac.Model{}
	err := enforcer.Allowed(model)
	assert.Nil(t, err)
}

func TestEnforcer_NotAllowed(t *testing.T) {
	loader := &MockLoader{
		Allowed: false,
	}
	enforcer := rbac.NewEnforcer(loader)
	model := &rbac.Model{}
	err := enforcer.Allowed(model)
	assert.NotNil(t, err)
}
