package accesscontrol

import "github.com/casbin/casbin/v2"

type AccessChecker struct {
	enforcer *casbin.Enforcer
}

func NewChecker(enforcer *casbin.Enforcer) *AccessChecker {
	return &AccessChecker{enforcer: enforcer}
}

func (ac *AccessChecker) CheckAccess(sub interface{}, obj, act string) (bool, error) {
	switch s := sub.(type) {
	case string:
		// RBAC using userID string
		return ac.enforcer.Enforce(s, obj, act)
	default:
		// fallback to ABAC: struct with attribute (eg. IsOwner)
		return ac.enforcer.Enforce(sub, obj, act)
	}
}
