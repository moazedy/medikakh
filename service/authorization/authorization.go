package authorization

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/casbin/casbin"
)

func IsPermissioned(role, object, action string) bool {
	baseFileName := "basic_model.conf"
	policyFileName := "basic_policy.csv"
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	policyPath := basePath + "/" + baseFileName
	modelPath := basePath + "/" + policyFileName

	enforcer, err := casbin.NewEnforcerSafe(modelPath, policyPath)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if enforcer == nil {
		log.Println("failed to open casbin files")
		return false
	}
	res := enforcer.Enforce(role, object, action)

	if !res {
		log.Println("access denied")
		return res
	}

	return res
}
