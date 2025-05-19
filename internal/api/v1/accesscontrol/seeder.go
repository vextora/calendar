package accesscontrol

import casbin "github.com/casbin/casbin/v2"

func SeedPolicy(e *casbin.Enforcer) {
	e.AddPolicy("editor", "article", "create")
	e.AddPolicy("editor", "article", "update")
	e.AddPolicy("viewer", "article", "read")
}
