package accesscontrol

import (
	"net/http"
	"oncomapi/internal/api/v1/accesscontrol"
	ac "oncomapi/pkg/accesscontrol"
)

func ACProtect(action string, getResource func(r *http.Request) (string, map[string]any)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := accesscontrol.User{
				ID:    r.Header.Get("X-User-ID"),
				Role:  r.Header.Get("X-User-Role"),
				IsVIP: r.Header.Get("X-User-VIP") == "true",
			}

			resource, ctx := getResource(r)
			req := accesscontrol.MakeAccessRequest(user, action, resource, ctx)
			if !ac.CheckAccess(req) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
