package rbac

type Enforcer interface {
	//Allowed checks if subject can act upon a resource
	Allowed(*Model) error
}
