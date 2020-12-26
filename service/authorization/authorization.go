package authorization

import (
	"log"

	"github.com/casbin/casbin"
)

func IsPremissioned(role, object, action string) bool {
	enforcer := casbin.NewEnforcer("auth_model.conf", "policy.csv")
	res := enforcer.Enforce(role, object, action)

	if !res {
		log.Println("access denied")
		return res
	}

	return res
}
