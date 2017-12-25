package pmapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

//查询出所有变量的内容
func (pmapi *PMApi) ListPsVariables(c *gin.Context) {
	var aryvars []proxysql.Variables

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")
	/*
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		if limit == 0 {
			limit = 10
		}

		if page == 0 {
			page = 1
		}

		skip := (page - 1) * limit
	*/

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.Variables{})
	} else {
		pmapi.PMhost = hostname
		pmapi.PMport, _ = strconv.ParseUint(port, 10, 64)
		pmapi.PMuser = username
		pmapi.PMpass = password

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		aryvars, pmapi.ApiErr = proxysql.GetConfig(pmapi.Apidb)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}
		c.JSON(http.StatusOK, aryvars)
	}

}

func (pmapi *PMApi) UpdateOneVariables(c *gin.Context) {

	var tmpvars proxysql.Variables

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.Variables{})
	} else {
		pmapi.PMhost = hostname + ":" + port
		pmapi.PMuser = username
		pmapi.PMpass = password

		pmapi.PMconn, pmapi.ApiErr = proxysql.NewConn(pmapi.PMhost, pmapi.PMport, pmapi.PMuser, pmapi.PMpass)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}

		pmapi.Apidb, pmapi.ApiErr = pmapi.PMconn.OpenConn()
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": pmapi.ApiErr})
		}
		if pmapi.ApiErr = c.Bind(&tmpvars); pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": pmapi.ApiErr})
		}

		pmapi.ApiErr = proxysql.UpdateOneConfig(pmapi.Apidb, tmpvars.VariablesName, tmpvars.Value)
		if pmapi.ApiErr != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": "UpdateOneVariable Failed"})
		}
		c.JSON(http.StatusOK, gin.H{"result": "OK"})
	}

}
