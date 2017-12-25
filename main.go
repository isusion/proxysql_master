package main

import (
	"github.com/imSQL/proxysql_master/pmapi"
)

func main() {

	// New api instance
	pmapiv1 := new(pmapi.PMApi)

	pmapiv1.RegisterServices()

	pmapiv1.RunApiService()

}
