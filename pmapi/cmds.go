package pmapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

func (pmapi *PMApi) SetProxySQLReadonly(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyReadOnly(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLReadwrite(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyReadWrite(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLStart(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyStart(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLRestart(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyRestart(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLStop(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyStop(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLPause(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyPause(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLResume(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyResume(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLShutdown(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyShutdown(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLFlogs(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyFlushLogs(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLKill(c *gin.Context) {
	pmapi.ApiErr = proxysql.ProxyKill(pmapi.Apidb)
	if pmapi.ApiErr != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
