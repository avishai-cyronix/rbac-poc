package rbac

type Factory interface {
	//Get policy enforcer
	Get(string, string) (Loader, error)
}

func NewFactory() Factory {
	return &factory{}
}

type factory struct{}

//Get policy enforcer
func (f *factory) Get(modelPath, policyPath string) (Loader, error) {
	return NewLoader(modelPath, policyPath)
}
