package accesscontrol

func checkRBAC(subject, action, object string) bool {
	return subject == "admin"
}

func checkABAC(req AccessRequest) bool {
	if isOwner, ok := req.Context["IsOwner"].(bool); ok && isOwner {
		return true
	}
	return false
}

func checkCasbin(subject, action, object string) bool {
	if GlobalChecker == nil {
		return false
	}
	ok, err := GlobalChecker.CheckAccess(subject, object, action)
	if err != nil {
		return false
	}
	return ok
}

func CheckAccess(req AccessRequest) bool {
	if checkRBAC(req.Subject, req.Action, req.Object) {
		return true
	}

	if checkABAC(req) {
		return true
	}

	if checkCasbin(req.Subject, req.Action, req.Object) {
		return true
	}
	return false
}
