package pmapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

func (pmapi *PMApi) SetProxySQLReadonly(c *gin.Context) {
	_, err := proxysql.ProxyReadOnly(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLReadwrite(c *gin.Context) {
	_, err := proxysql.ProxyReadWrite(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLStart(c *gin.Context) {
	_, err := proxysql.ProxyStart(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLRestart(c *gin.Context) {
	_, err := proxysql.ProxyRestart(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLStop(c *gin.Context) {
	_, err := proxysql.ProxyStop(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLPause(c *gin.Context) {
	_, err := proxysql.ProxyPause(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLResume(c *gin.Context) {
	_, err := proxysql.ProxyResume(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLShutdown(c *gin.Context) {
	_, err := proxysql.ProxyShutdown(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLFlogs(c *gin.Context) {
	_, err := proxysql.ProxyFlushLogs(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}

func (pmapi *PMApi) SetProxySQLKill(c *gin.Context) {
	_, err := proxysql.ProxyKill(pmapi.Apidb)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
