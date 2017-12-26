package main

import (
	"fmt"
	"os"

	"github.com/imSQL/proxysql_master/pmapi"
	"github.com/juju/errors"
)

func main() {

	// New api instance
	pmapiv1, err := pmapi.NewApi()
	if err != nil {
		fmt.Fprintf(os.Stdout, "pmapi instance initialization failure")
		errors.Details(err)
		os.Exit(1)
	}

	// set api runnings args.
	pmapiv1.SetApiLog(os.Stdout)
	pmapiv1.SetApiPort(3333)

	// register router.
	pmapiv1.RegisterServices()

	// running.
	pmapiv1.RunApiService()

}
