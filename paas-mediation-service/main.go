package main

//go:generate go get github.com/swaggo/swag/cmd/swag@v1.16.3
//go:generate go run github.com/swaggo/swag/cmd/swag init --generalInfo /main.go --parseDependency --parseDepth 2

import (
	"github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/lib"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// swagger docs
	_ "github.com/netcracker/qubership-core-paas-mediation/paas-mediation-service/v2/docs"

	// memlimit sets memory limit = 0.9 of cgroup memory limit
	_ "github.com/netcracker/qubership-core-lib-go/v3/memlimit"
)

// @title           			Paas Mediation API
// @version         			2.0
// @description     			API for Paas Mediation.
// @tag.name                    api version info
// @tag.description             Apis provides information related to versions
// @tag.name                    since:2.0
// @tag.description             Apis existed since 2.0 version
// @Produce 					json
// @securityDefinitions.apikey 	ApiKeyAuth
// @in 							header
// @name 						Authorization
func main() {
	lib.RunServer()
}
