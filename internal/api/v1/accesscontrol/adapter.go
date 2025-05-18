package accesscontrol

import ac "oncomapi/pkg/accesscontrol"

type User struct {
	ID    string
	Role  string
	IsVIP bool
}

func MakeAccessRequest(user User, action, resource string, ctx map[string]any) ac.AccessRequest {
	ctx["is_vip"] = user.IsVIP
	return ac.AccessRequest{
		Subject:  user.Role,
		Action:   action,
		Resource: resource,
		Context:  ctx,
	}
}
