package pmapi

import (
	"database/sql"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/imSQL/proxysql"
)

type PMApi struct {
	// proxysql connection informations.
	PMuser string
	PMpass string
	PMhost string
	PMport uint64
	// proxysql connnection handler
	PMconn *proxysql.Conn
	// sql database handler
	Apidb *sql.DB
	// restfulapi log
	ApiLog io.Writer
	// restfulapi error
	ApiErr error
	// gin engine
	Router *gin.Engine
	// api port
	ApiPort int
}

// new api instance.
func NewApi() (*PMApi, error) {
	pmapi := new(PMApi)

	pmapi.PMhost = "127.0.0.1"
	pmapi.PMport = 6032
	pmapi.PMuser = "admin"
	pmapi.PMpass = "admin"

	pmapi.ApiLog = os.Stdout
	pmapi.ApiPort = 3333

	return pmapi, nil
}

// set api log.
func (pmapi *PMApi) SetApiLog(w io.Writer) {
	pmapi.ApiLog = w
}

// set api port.
func (pmapi *PMApi) SetApiPort(port int) {
	pmapi.ApiPort = port
}

func (pmapi *PMApi) RegisterServices() {

	/*初始化gin实例*/
	pmapi.Router = gin.Default()

	/*Dashboard*/
	pmapi.Router.GET("/api/v1/status", pmapi.ListPStatus)

	/*Variables*/
	pmapi.Router.GET("/api/v1/variables", pmapi.ListPsVariables)
	pmapi.Router.PUT("/api/v1/variables", pmapi.UpdateOneVariables)

	/*User Services*/
	pmapi.Router.GET("/api/v1/users", pmapi.ListAllUsers)
	pmapi.Router.POST("/api/v1/users", pmapi.CreateOneUser)
	pmapi.Router.PUT("/api/v1/users", pmapi.UpdateOneUser)
	pmapi.Router.DELETE("/api/v1/users", pmapi.DeleteOneUser)

	/*Server Services*/
	pmapi.Router.GET("/api/v1/servers", pmapi.ListAllServers)
	pmapi.Router.POST("/api/v1/servers", pmapi.CreateOneServer)
	pmapi.Router.PUT("/api/v1/servers", pmapi.UpdateOneServer)
	pmapi.Router.DELETE("/api/v1/servers", pmapi.DeleteOneServers)

	/*Query Rules*/
	pmapi.Router.GET("/api/v1/queryrules", pmapi.ListAllQueryRules)
	pmapi.Router.POST("/api/v1/queryrules", pmapi.CreateOneQueryRules)
	pmapi.Router.PUT("/api/v1/queryrules", pmapi.UpdateOneQueryRules)
	pmapi.Router.DELETE("/api/v1/queryrules", pmapi.DeleteOneQueryRules)

	/*Scheduler*/
	pmapi.Router.GET("/api/v1/schedulers", pmapi.ListAllScheduler)
	pmapi.Router.POST("/api/v1/schedulers", pmapi.CreateOneScheduler)
	pmapi.Router.PUT("/api/v1/schedulers", pmapi.UpdateOneScheduler)
	pmapi.Router.DELETE("/api/v1/schedulers", pmapi.DeleteOneScheduler)

	/*ProxySQL admin API*/
	pmapi.Router.GET("/api/v1/cmd/readonly", pmapi.SetProxySQLReadonly)
	pmapi.Router.GET("/api/v1/cmd/readwrite", pmapi.SetProxySQLReadwrite)
	pmapi.Router.GET("/api/v1/cmd/start", pmapi.SetProxySQLStart)
	pmapi.Router.GET("/api/v1/cmd/restart", pmapi.SetProxySQLRestart)
	pmapi.Router.GET("/api/v1/cmd/stop", pmapi.SetProxySQLStop)
	pmapi.Router.GET("/api/v1/cmd/pause", pmapi.SetProxySQLPause)
	pmapi.Router.GET("/api/v1/cmd/resume", pmapi.SetProxySQLResume)
	pmapi.Router.GET("/api/v1/cmd/shutdown", pmapi.SetProxySQLShutdown)
	pmapi.Router.GET("/api/v1/cmd/flushlogs", pmapi.SetProxySQLFlogs)
	pmapi.Router.GET("/api/v1/cmd/kill", pmapi.SetProxySQLKill)

}

func (pmapi *PMApi) RunApiService() {

	pmapi.Router.Run(":" + strconv.Itoa(pmapi.ApiPort))
}
