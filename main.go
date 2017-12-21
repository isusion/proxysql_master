package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/imSQL/proxysql-master/pmapi"
)

var (
	apiSource = flag.String("s", "admin/admin@localhost:6032/main", "ProxySQL Connection URI address.")
	apiPort   = flag.Int64("p", 6031, "Api port.")
	apiLog    = flag.String("l", "/tmp/pm.log", "api log file.")
)

func main() {

	// New api instance
	pmapiv1 := new(pmapi.PMApi)
	pmapiv1.ApiHost = fmt.Sprintf(":%d", *apiPort)

	pmapiv1.PMuser = strings.Split(strings.Split(*apiSource, "@")[0], "/")[0]
	pmapiv1.PMpass = strings.Split(strings.Split(*apiSource, "@")[0], "/")[1]
	pmapiv1.PMhost = strings.Split(strings.Split(*apiSource, "@")[1], "/")[0]
	pmapiv1.PMdb = strings.Split(strings.Split(*apiSource, "@")[1], "/")[1]

	pmapiv1.MakePMdbi()

	pmapiv1.RegisterServices()

	pmapiv1.RunApiService()

}
