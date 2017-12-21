package pmapi

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imSQL/proxysql"
)

/*返回所有后端数据库服务器的信息*/
func (pmapi *PMApi) ListAllServers(c *gin.Context) {

	var tmpserver proxysql.Servers
	var aryservers []proxysql.Servers

	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	skip := (page - 1) * limit

	if len(hostname) == 0 || hostname == "undefined" {
		c.JSON(http.StatusOK, []proxysql.Servers{})
	} else {
		pmapi.PMhost = hostname + ":" + port
		pmapi.PMuser = username
		pmapi.PMpass = password
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		defer pmapi.Apidb.Close()

		aryservers, err = tmpserver.FindAllServerInfo(pmapi.Apidb, limit, skip)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, aryservers)
	}

}

/*创建一个新的后端数据库服务节点*/
func (pmapi *PMApi) CreateOneServer(c *gin.Context) {

	var tmpserver proxysql.Servers
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.Servers{})
	} else {
		pmapi.PMhost = hostname + ":" + port
		pmapi.PMuser = username
		pmapi.PMpass = password
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		defer pmapi.Apidb.Close()

		if err := c.Bind(&tmpserver); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->CreateOneServer->AddOneServer tmpserver", tmpserver)

		_, err := tmpserver.AddOneServers(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->CreateOneServer->AddOneServer Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

/*删除指定服务器*/
func (pmapi *PMApi) DeleteOneServers(c *gin.Context) {
	var tmpserver proxysql.Servers
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.Servers{})
	} else {
		pmapi.PMhost = hostname + ":" + port
		pmapi.PMuser = username
		pmapi.PMpass = password
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		defer pmapi.Apidb.Close()

		if err := c.Bind(&tmpserver); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->DeleteOneServer->DeleteOneServer tmpserver", tmpserver)

		_, err := tmpserver.DeleteOneServers(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->DeleteOneServer->DeleteOneServer Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}

/*更新服务信息的patch函数*/
func (pmapi *PMApi) UpdateOneServer(c *gin.Context) {
	var tmpserver proxysql.Servers
	var err error

	hostname := c.Query("hostname")
	port := c.Query("port")
	username := c.Query("adminuser")
	password := c.Query("adminpass")

	if len(hostname) == 0 {
		c.JSON(http.StatusOK, []proxysql.Servers{})
	} else {
		pmapi.PMhost = hostname + ":" + port
		pmapi.PMuser = username
		pmapi.PMpass = password
		pmapi.PMdb = "information_schema"
		pmapi.MakePMdbi()

		pmapi.Apidb, err = sql.Open("mysql", pmapi.PMdbi)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"error": err})
		}
		defer pmapi.Apidb.Close()

		if err := c.Bind(&tmpserver); err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		}
		log.Print("pmapi->UpdateOneServer->UpdateOneServer tmpserver", tmpserver)

		_, err := tmpserver.UpdateOneServerInfo(pmapi.Apidb)
		if err != nil {
			log.Print("pmapi->UpdateOneServer->UpdateOneServer Failed", err)
			c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "OK"})
		}
	}
}
