package accesscontrol

import (
	"log"

	casbin "github.com/casbin/casbin/v2"
)

func SeedPolicy(e *casbin.Enforcer) {
	if _, err := e.AddPolicy("editor", "article", "create"); err != nil {
		log.Println("AddPolicy error:", err)
	}

	if _, err := e.AddPolicy("editor", "article", "update"); err != nil {
		log.Println("AddPolicy error:", err)
	}

	if _, err := e.AddPolicy("viewer", "article", "read"); err != nil {
		log.Println("AddPolicy error:", err)
	}
}
