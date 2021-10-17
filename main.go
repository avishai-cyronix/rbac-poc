package main

import (
	"my-module/rbac"
)

func main() {
	//TODO: Use mongo adapter to save policies in the db instead csv file.
	//TODO: https://github.com/casbin/mongodb-adapter

	fac := rbac.NewFactory()
	modelPath := "config/rbac_model.conf"
	policyPath := "config/rbac_policy.csv"
	loader, err := fac.Get(modelPath, policyPath)
	if err != nil {
		return
	}

	enforcer := rbac.NewEnforcer(loader)
	model := &rbac.Model{
		Role:   "project_owner",
		Object: "project",
		Act:    "write",
	}

	okErr := enforcer.Allowed(model)
	if okErr != nil {
		println("PERMISSION DENIED!")
	} else {
		println("PERMISSION GRANTED!")
	}
}
