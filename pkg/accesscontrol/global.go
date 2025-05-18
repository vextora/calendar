package accesscontrol

import "github.com/casbin/casbin/v2"

var GlobalChecker *AccessChecker

func InitGlobal(enforcer *casbin.Enforcer) {
	GlobalChecker = NewChecker(enforcer)
}
